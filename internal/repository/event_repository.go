package repository

import (
	"Ereno_community_event_system/internal/config"
	"Ereno_community_event_system/internal/models"
)

func GetAllEvents() ([]models.Event, error) {
	rows, err := config.DB.Query("SELECT id, name, date, time, venue, description FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		rows.Scan(&e.ID, &e.Name, &e.Date, &e.Time, &e.Venue, &e.Description)
		events = append(events, e)
	}
	return events, nil
}

func AddEvent(e models.Event) error {
	_, err := config.DB.Exec("INSERT INTO events (name, date, time, venue, description) VALUES (?, ?, ?, ?, ?)",
		e.Name, e.Date, e.Time, e.Venue, e.Description)
	return err
}

// ✅ Add this function to fetch one event by ID
func GetEventByID(id int) (models.Event, error) {
	row := config.DB.QueryRow("SELECT id, name, date, time, venue, description FROM events WHERE id = ?", id)

	var e models.Event
	err := row.Scan(&e.ID, &e.Name, &e.Date, &e.Time, &e.Venue, &e.Description)
	return e, err
}
