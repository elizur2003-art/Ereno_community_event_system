package repository

import (
	"Ereno_community_event_system/internal/config"
	"Ereno_community_event_system/internal/models"
)

func AddFeedback(f models.Feedback) error {
	_, err := config.DB.Exec(
		"INSERT INTO feedback (event_id, name, comment, rating) VALUES (?, ?, ?, ?)",
		f.EventID, f.Name, f.Comment, f.Rating,
	)
	return err
}

func GetFeedbackByEvent(eventID int) ([]models.Feedback, error) {
	rows, err := config.DB.Query(
		"SELECT id, event_id, name, comment, rating FROM feedback WHERE event_id = ?",
		eventID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []models.Feedback
	for rows.Next() {
		var f models.Feedback
		rows.Scan(&f.ID, &f.EventID, &f.Name, &f.Comment, &f.Rating)
		feedbacks = append(feedbacks, f)
	}
	return feedbacks, nil
}
