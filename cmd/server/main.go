package main

import (
	"Ereno_community_event_system/internal/config"
	"Ereno_community_event_system/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	config.LoadEnv()
	config.InitDB()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/events", handlers.EventListHandler)
	mux.HandleFunc("/events/create", handlers.CreateEventHandler)
	http.HandleFunc("/events/store", handlers.HomeHandle)
	mux.HandleFunc("/attendees/register", handlers.RegisterAttendeeHandler)
	mux.HandleFunc("/reports", handlers.ReportHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/feedback", handlers.FeedbackHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}
	fmt.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
