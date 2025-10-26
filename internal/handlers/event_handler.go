package handlers

import (
	"Ereno_community_event_system/internal/models"
	"Ereno_community_event_system/internal/repository"
	"html/template"
	"net/http"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/events", http.StatusSeeOther)
}

func EventListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/events.html"))
	events, _ := repository.GetAllEvents()
	tmpl.Execute(w, events)
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		event := models.Event{
			Name:        r.FormValue("name"),
			Date:        r.FormValue("date"),
			Time:        r.FormValue("time"),
			Venue:       r.FormValue("venue"),
			Description: r.FormValue("description"),
		}
		repository.AddEvent(event)
		http.Redirect(w, r, "/events", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("web/templates/create_event.html"))
	tmpl.Execute(w, nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	tmpl.Execute(w, nil)
}
