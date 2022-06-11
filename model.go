package main

import "time"

//structure for records response
type Record struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks int64     `json:"totalMarks"`
}

//structure for response json
type Response struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Records []Record `json:"records"`
}

//structure for request payload
type Request struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}
