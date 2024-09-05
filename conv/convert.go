package conv

import (
	"fmt"
	"or/http"
	"or/util"
	"strconv"
	"strings"
	"time"
)

var utcOffset int

func ToUTC(fromTime time.Time, cityCode string) time.Time {
	// Create a Duration representing the UTC offset in seconds

	offsetDuration := time.Duration(utcOffset) * time.Second

	// Subtract the offset from the local time to get UTC time
	utcTime := fromTime.Add(-offsetDuration)

	return utcTime
}

// provide the UTC time and the city code and it will return the local time
func ToLocalTime(utcTime time.Time, cityCode string) time.Time {
	// Create a Duration representing the UTC offset in seconds
	offsetDuration := time.Duration(utcOffset) * time.Second

	// Add the offset to the UTC time to get local time
	localTime := utcTime.Add(offsetDuration)

	return localTime
}

// provide the UTC time and the city code and it will return the local time
func ToGDSDate(convertDate time.Time, includeYear bool) string {
	var formattedDate string
	if includeYear {
		// Convert the date to "ddMMM" format with no year
		formattedDate = convertDate.Format("02Jan")

	} else {
		// Convert the date to "ddMMMyy" format with no year
		formattedDate = convertDate.Format("02Jan24")

	}
	return formattedDate
}

/*
	provide the UTC time and the city code and it will return the local time

dateformat: indicate like this where day month and year are in the string
use like this: m/d/y
*/
func DateFromString(convertDateString string, timeString string, dateFormat string) time.Time {
	dateFormat = strings.ToUpper(dateFormat)

	matches := util.RegexGroups(`(?P<val1>\w+)[-/\,\s](?P<val2>\w+)[-/\,\s](?P<val3>\w+)`, convertDateString)
	format := util.RegexGroups(`(?P<fmt1>[A-Z])[-/\,\s](?P<fmt2>[A-Z])[-/\,\s](?P<fmt3>[A-Z])`, dateFormat)

	if matches != nil {
		var day string
		var month string
		var year string
		var layout = "2006-01-02"

		// assign the month first
		if format["fmt1"] == "M" {
			month = matches["val1"]
		} else if format["fmt2"] == "M" {
			month = matches["val2"]
		} else {
			month = matches["val3"]
		}

		if format["fmt1"] == "D" {
			day = matches["val1"]
		} else if format["fmt2"] == "D" {
			day = matches["val2"]
		} else {
			day = matches["val3"]
		}

		if format["fmt1"] == "Y" {
			year = matches["val1"]
		} else if format["fmt2"] == "Y" {
			year = matches["val2"]
		} else {
			year = matches["val3"]
		}

		if len(year) < 4 {
			year += "20" + year
		}
		if len(month) < 2 {
			month = "0" + month
		}
		if len(day) < 2 {
			day = "0" + day
		}
		convertDateString := year + "-" + month + "-" + day
		if timeString != "" {
			times := util.RegexGroups(`(?i)(?P<tim1>\d{1,2})[-/\,\s:]*(?P<tim2>\d{2})[-/\,\s]*(?P<tim3>(A|P))?`, timeString)
			//if they sent in with "P" or "PM", needs to be converted to 24 hour
			if strings.ToUpper(times["tim3"]) == "P" {
				timeAsNumber, _ := strconv.Atoi(times["tim1"])
				if timeAsNumber < 12 {
					timeAsNumber += 12
					times["tim1"] = strconv.Itoa(timeAsNumber)
				}
			}

			if len(times["tim1"]) < 2 {
				times["tim1"] = "0" + times["tim1"]
			}
			timeString = times["tim1"] + ":" + times["tim2"]
			layout += " 15:04"
			convertDateString += " " + timeString
		}

		date, err := time.Parse(layout, convertDateString)
		if err != nil {
			// Handle parsing error
			fmt.Println("Error parsing date:", err)
			return time.Time{}
		}
		return date
	}
	return time.Time{}
}

///// CONVERT CURRENCY //////

func Currency(from string, to string) float64 {

	// define what is required in the return response struct
	type Data struct {
		Result         string  `json:"result"`
		Documentation  string  `json:"documentation"`
		ConversionRate float64 `json:"conversion_rate"`
	}
	data := Data{}

	var rate float64
	headers := map[string]string{}
	urlString := "https://v6.exchangerate-api.com/v6/YOUR-API-KEY/pair/EUR/GBP"
	results, err := http.Get(urlString, headers, data, "")
	if err != nil {
		if results.StatusCode == 200 {
			type Response struct {
				Result         string  `json:"result"`
				ConversionRate float64 `json:"conversion_rate"`
			}
			response := Response{}
			return responseResponse.ConversionRate

		}

	}
	return rate

}
