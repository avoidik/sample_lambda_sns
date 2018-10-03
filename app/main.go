package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type eventData struct {
	Operation string `json:"operation"`
	ID        int    `json:"id"`
	Data      string `json:"data"`
}

func (e *eventData) sync() {
	switch e.Operation {
	case "CREATE":
		e.createOrUpdateDB()
	case "UPDATE":
		e.createOrUpdateDB()
	case "DELETE":
		e.deleteDB()
	default:
		fmt.Println("Unknown operation", e.Operation)
	}
}

func handler(event events.SNSEvent) {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		fmt.Println("Error during decoding event data")
		return
	}
	fmt.Println("SNS event:", string(eventJSON))

	for _, record := range event.Records {
		msg := record.SNS.Message
		subj := record.SNS.Subject
		fmt.Println("Message", msg, "Subject", subj)
		e := &eventData{}
		json.Unmarshal([]byte(msg), e)
		e.sync()
	}
}

func main() {
	lambda.Start(handler)
}
