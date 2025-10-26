package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ No .env file found — using system environment variables")
	}
}

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./community.db")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	createTables()
	fmt.Println("✅ Database connected successfully")
}

func createTables() {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			date TEXT,
			time TEXT,
			venue TEXT,
			description TEXT
		);`,

		`CREATE TABLE IF NOT EXISTS attendees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			event_id INTEGER,
			FOREIGN KEY (event_id) REFERENCES events(id)
		);`,

		`CREATE TABLE IF NOT EXISTS feedback (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			name TEXT,
			comment TEXT,
			rating INTEGER,
			FOREIGN KEY (event_id) REFERENCES events(id)
		);`,
	}

	for _, table := range tables {
		_, err := DB.Exec(table)
		if err != nil {
			log.Fatal("Error creating table:", err)
		}
	}
}
