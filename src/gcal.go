package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"encoding/json"
	"fmt"
)

const LIST_CALENDAR_EVENTS = "https://www.googleapis.com/calendar/v3/calendars/calendarId/events"
const LIST_CALENDARS = "https://www.googleapis.com/calendar/v3/calendars"

type Calendar struct {
	Id   string `json:"id"`
	Name string `json:"summary"`
}

type GoogleCalendarCredentials struct {
	APIKey string
}

// just wanted to try generics
func marshal[T any](parsed []T) ([]byte, error) {
	return json.Marshal(parsed)
}

func getAllRooms() []Calendar {

	// Make a GET request to the calendar API with the specified time range
	resp, err := http.Get(LIST_CALENDARS)
	if err != nil {
		// Return an empty slice if an error occurs during the GET request
		return make([]Calendar, 0)
	}
	defer resp.Body.Close()

	// Read the response body and parse the events
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]Calendar, 0)
	}
	return parseCalendars(body)

}
func getCalendarById(calendars []Calendar, id string) Calendar {
	for _, c := range calendars {
		if c.Id == id {
			return c
		}
	}
	return Calendar{}
}

// parseCalendars parses the given JSON data into a slice of Calendar structs.
// It takes a byte slice of JSON data as input and returns a slice of Calendar structs.
// If there is an error parsing the JSON data, it will print an error message and return an empty slice of Calendar structs.
func parseCalendars(data []byte) []Calendar {
	var calendars struct {
		Items []Calendar `json:"items"`
	}
	err := json.Unmarshal(data, &calendars)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return make([]Calendar, 0)
	}

	var parsedCalendars []Calendar
	for _, obj := range calendars.Items {
		if strings.Contains(obj.Id, "resource.calendar.google") {
			parsedCalendars = append(parsedCalendars, Calendar{Name: obj.Name, Id: obj.Id})
		}
	}
	return parsedCalendars
}

// getRoomEvents retrieves events for a specific room on a given day.
// It takes in the name of the room and the day in the format "YYYY-MM-DD" as parameters.
// It returns a slice of ParsedEvent structs representing the events for the room on that day.
// If there are no events or an error occurs during the retrieval process, an empty slice is returned.
func getRoomEvents(calendarId string, day string) []ParsedEvent {

	// Get the minimum and maximum dates for the given day
	timeMin, timeMax := getMinMaxDates(day)

	// Make a GET request to the calendar API with the specified time range
	resp, err := http.Get(getCalendarURL(calendarId) + fmt.Sprintf("?timeMin=%s&timeMax=%s", timeMin, timeMax))
	if err != nil {
		// Return an empty slice if an error occurs during the GET request
		return make([]ParsedEvent, 0)
	}
	defer resp.Body.Close()

	// Read the response body and parse the events
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return make([]ParsedEvent, 0)
	}
	return parseEvents(body)

	// Return an empty slice if the room calendar ID is not found
	return make([]ParsedEvent, 0)
}

func getCalendarURL(calendarId string) string {
	return strings.Replace(LIST_CALENDAR_EVENTS, "calendarId", calendarId, -1)
}

func NewGoogleCalendarCredentials() GoogleCalendarCredentials {
	return GoogleCalendarCredentials{
		APIKey: os.Getenv("GOOGLE_CALENDAR_API_KEY"),
	}
}
