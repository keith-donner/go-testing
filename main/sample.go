package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"or/debug"
	"or/o"
	"or/rmk"
	"strings"
	"time"
)

type Queue struct {
	List []struct {
		Category string `json:"category"`
		Pcc      string `json:"pcc"`
		Number   int    `json:"number"`
	} `json:"list"`
	Agent string    `json:"agent"`
	Time  time.Time `json:"time"`
}

type QueueHistoryList struct {
	PnrID  string  `json:"pnrID"`
	Ttl    int64   `json:"ttl"`
	Queues []Queue `json:"queues"`
}

/////////////////////////////////////////////////////////

func sample() {
	// this is the structure of the json response

	pnr := o.Pnr{}
	remark := pnr.Remark
	var (
		jsonResponse string
		compareDate  time.Time
		queueNumber  int
		params       map[string]string
		pcc          string
		agent        string
	)
	// Get current date
	compareDate = time.Now()

	// Get Queue History from AWS
	url := fmt.Sprintf("https://xgd6mf0gzk.execute-api.us-east-2.amazonaws.com/prod/qh?pnrID=%s", pnr.Id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching queue history:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	jsonResponse = string(body)

	if strings.Contains(jsonResponse, "queues") {
		// Parse JSON into queueHistoryList
		var queueList QueueHistoryList

		err = json.Unmarshal([]byte(jsonResponse), queueList)

		for _, queue := range queueList.Queues {
			if queue.Time.Before(compareDate) {
				for _, listItem := range queue.List {
					if listItem.Number == 90 {
						// Find existing remark
						remark = rmk.Find(o.GeneralRemark, "QUEUETO-")
						if remark.Contents != "" {
							// No existing remark, create new one
							remarkText := "QUEUETO-" + pcc + "/" + string(queueNumber) + " BY AGENT-" + agent

							// Add remark to PNR
							params["gds"] = "a"
							params["cat"] = "v"
							rmk.Add(remarkText, o.GeneralRemark, params, []int{}, []int{})
							debug.LogEmail("ALERT: NEW agent queued PNR to queue 90: "+pnr.This.OwningOffice.ID+"- rloc: "+pnr.This.RecordLocator, "check the PNR", "keith.donner@business-class.com")
						}
					}
				}
			}
		}
	}
}
