package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Creator struct {
		Email string `json:"email"`
	} `json:"creator"`
	Start struct {
		DateTime string `json:"dateTime"`
	} `json:"start"`
	End struct {
		DateTime string `json:"dateTime"`
	} `json:"end"`
}

type ParsedEvent struct {
	Creator string
	Start   time.Time
	End     time.Time
}

func parseEvents(data []byte) []ParsedEvent {
	var events struct {
		Items []Event `json:"items"`
	}
	err := json.Unmarshal(data, &events)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return make([]ParsedEvent, 0)
	}

	var parsedEvents []ParsedEvent
	for _, event := range events.Items {
		fmt.Printf("Event booked by: %s\n", event.Creator.Email)
		startTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
		if err != nil {
			fmt.Println("Error parsing start time:", err)
			continue
		}
		endTime, err := time.Parse(time.RFC3339, event.End.DateTime)
		if err != nil {
			fmt.Println("Error parsing end time:", err)
			continue
		}
		fmt.Printf("Start time: %s\n", startTime.Format(time.RFC3339))
		fmt.Printf("End time: %s\n", endTime.Format(time.RFC3339))
		parsedEvents = append(parsedEvents, ParsedEvent{Creator: event.Creator.Email, Start: startTime, End: endTime})
	}
	return parsedEvents
}
