package services

import (
	"Ereno_community_event_system/internal/models"
	"Ereno_community_event_system/internal/repository"
)

type EventReport struct {
	Event       models.Event
	Attendees   []models.Attendee
	Feedbacks   []models.Feedback
	TotalGuests int
	AvgRating   float64
}

func GenerateEventReport(eventID int) (EventReport, error) {
	var report EventReport

	// Get Event Info
	events, _ := repository.GetAllEvents()
	for _, e := range events {
		if e.ID == eventID {
			report.Event = e
			break
		}
	}

	// Get Attendees
	attendees, _ := repository.GetAttendeesByEvent(eventID)
	report.Attendees = attendees
	report.TotalGuests = len(attendees)

	// Get Feedback
	feedbacks, _ := repository.GetFeedbackByEvent(eventID)
	report.Feedbacks = feedbacks

	// Compute average rating
	var totalRating int
	for _, f := range feedbacks {
		totalRating += f.Rating
	}
	if len(feedbacks) > 0 {
		report.AvgRating = float64(totalRating) / float64(len(feedbacks))
	}

	return report, nil
}
