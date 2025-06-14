package models

import (
	"go-dev/module/db"
	"time"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `default:"now" time_format:"2006-01-02 15:04:05"`
	UserID      int
}

func (e Event) Save() {
	query := "INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		panic(err)
	}
	id, error := result.LastInsertId()
	if error != nil {
		panic(error)
	}
	e.ID = int(id)
}
func GetEvents() []Event {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			panic(err)
		}
		events = append(events, event)
	}
	return events
}
