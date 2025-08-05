package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	Date        time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// TODO Add it to a database later
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
