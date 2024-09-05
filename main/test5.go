package main

// Define the structs to match the JSON structure
type List struct {
	Category string `json:"category"`
	Number   int    `json:"number"`
	PCC      string `json:"pcc"`
}

type Queue struct {
	Agent string `json:"agent"`
	List  []List `json:"list"`
	Time  string `json:"time"`
}

type PNRResponse struct {
	PNRID string  `json:"pnrID"`
	Queues []Queue `json:"queues"`
	TTL   int64   `json:"ttl"`
}



func test(){
	pnrData := PNRResponse{}
	for _, value := range pnrData.Queues {
		if   value.  {
			
	
	
	
		}
			
	
	}
	
	
	
}

