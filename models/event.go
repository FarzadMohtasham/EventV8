package models

import (
	"time"

	"github.com/FarzadMohtasham/EventV8/db"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.Date, e.UserID)

	if err != nil {
		return err
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = int(eventId)

	return nil
}

func GetAllEvents() (events []Event, err error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
