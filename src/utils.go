package main

import (
	"time"
)

func getMinMaxDates(day string) (string, string) {
	var timeMin, timeMax string
	switch day {
	case "today":
		timeMin = time.Now().Format("2006-01-02")
		timeMax = time.Now().AddDate(0, 0, 1).Format("2006-01-02T00:00:00Z")
	case "tomorrow":
		timeMin = time.Now().AddDate(0, 0, 1).Format("2006-01-02T00:00:00Z")
		timeMax = time.Now().AddDate(0, 0, 2).Format("2006-01-02T00:00:00Z")
	default:
		timeMin = time.Now().Format("2006-01-02T00:00:00Z")
		timeMax = time.Now().AddDate(0, 0, 1).Format("2006-01-02T00:00:00Z")
	}
	return timeMin, timeMax
}
