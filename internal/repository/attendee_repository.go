package repository

import (
	"Ereno_community_event_system/internal/config"
	"Ereno_community_event_system/internal/models"
)

// Add a new attendee to a specific event
func AddAttendee(a models.Attendee) error {
	_, err := config.DB.Exec(
		"INSERT INTO attendees (name, email, event_id) VALUES (?, ?, ?)",
		a.Name, a.Email, a.EventID,
	)
	return err
}

// Get all attendees for a specific event
func GetAttendeesByEvent(eventID int) ([]models.Attendee, error) {
	rows, err := config.DB.Query(
		"SELECT id, name, email, event_id FROM attendees WHERE event_id = ?",
		eventID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendees []models.Attendee
	for rows.Next() {
		var a models.Attendee
		rows.Scan(&a.ID, &a.Name, &a.Email, &a.EventID)
		attendees = append(attendees, a)
	}
	return attendees, nil
}

// Get all attendees (for reporting or admin view)
func GetAllAttendees() ([]models.Attendee, error) {
	rows, err := config.DB.Query("SELECT id, name, email, event_id FROM attendees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attendees []models.Attendee
	for rows.Next() {
		var a models.Attendee
		rows.Scan(&a.ID, &a.Name, &a.Email, &a.EventID)
		attendees = append(attendees, a)
	}
	return attendees, nil
}
