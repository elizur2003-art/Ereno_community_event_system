package handlers

import (
	"Ereno_community_event_system/internal/repository"
	"Ereno_community_event_system/internal/services"
	"html/template"
	"net/http"
	"strconv"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	// If user clicked a specific event to view report
	eventIDStr := r.URL.Query().Get("event_id")
	if eventIDStr != "" {
		eventID, _ := strconv.Atoi(eventIDStr)
		reportData, _ := services.GenerateEventReport(eventID)

		tmpl := template.Must(template.ParseFiles("web/templates/report_detail.html"))
		tmpl.Execute(w, reportData)
		return
	}

	// Otherwise, list all events for selection
	events, _ := repository.GetAllEvents()
	tmpl := template.Must(template.ParseFiles("web/templates/reports.html"))
	tmpl.Execute(w, events)

}
