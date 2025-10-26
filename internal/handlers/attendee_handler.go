package handlers

import (
	"Ereno_community_event_system/internal/models"
	"Ereno_community_event_system/internal/repository"
	"Ereno_community_event_system/internal/services"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func RegisterAttendeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		eventID, _ := strconv.Atoi(r.FormValue("event_id"))
		attendee := models.Attendee{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			EventID: eventID,
		}

		repository.AddAttendee(attendee)

		event, err := repository.GetEventByID(eventID)
		if err == nil {
			notificationService := services.NewNotificationService()
			message := fmt.Sprintf("You are successfully registered for the event: %s on %s at %s.",
				event.Name, event.Date, event.Venue)
			notificationService.SendNotification(attendee.Email, "Event Registration Confirmation", message)
		}

		http.Redirect(w, r, "/events", http.StatusSeeOther)
		return
	}

	events, _ := repository.GetAllEvents()
	tmpl := template.Must(template.ParseFiles("web/templates/register_attendee.html"))
	tmpl.Execute(w, events)
}
