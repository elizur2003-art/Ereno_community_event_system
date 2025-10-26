package handlers

import (
	"Ereno_community_event_system/internal/models"
	"Ereno_community_event_system/internal/repository"
	"html/template"
	"net/http"
	"strconv"
)

func FeedbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		eventID, _ := strconv.Atoi(r.FormValue("event_id"))
		feedback := models.Feedback{
			EventID: eventID,
			Name:    r.FormValue("name"),
			Comment: r.FormValue("comment"),
			Rating:  atoiDefault(r.FormValue("rating")),
		}
		repository.AddFeedback(feedback)
		http.Redirect(w, r, "/events", http.StatusSeeOther)
		return
	}

	eventIDStr := r.URL.Query().Get("event_id")
	eventID, _ := strconv.Atoi(eventIDStr)
	event, _ := repository.GetEventByID(eventID)

	tmpl := template.Must(template.ParseFiles("web/templates/feedback.html"))
	tmpl.Execute(w, event)
}

func atoiDefault(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
