package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// Load environment variables (works locally and on Render)
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️ No .env file found — using system environment variables")
	}
}

// Initialize PostgreSQL database connection
func InitDB() {
	var err error

	// Get full PostgreSQL connection string from Render
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("❌ DATABASE_URL environment variable not set")
	}

	// Connect to PostgreSQL using the connection URL
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL:", err)
	}

	// Test the database connection
	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Database unreachable:", err)
	}

	fmt.Println("✅ PostgreSQL database connected successfully")
	createTables()
}

// Create tables if they don’t exist
func createTables() {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT,
			date TEXT,
			time TEXT,
			venue TEXT,
			description TEXT
		);`,

		`CREATE TABLE IF NOT EXISTS attendees (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT,
			event_id INTEGER REFERENCES events(id)
		);`,

		`CREATE TABLE IF NOT EXISTS feedback (
			id SERIAL PRIMARY KEY,
			event_id INTEGER REFERENCES events(id),
			name TEXT,
			comment TEXT,
			rating INTEGER
		);`,
	}

	for _, table := range tables {
		_, err := DB.Exec(table)
		if err != nil {
			log.Fatal("❌ Error creating table:", err)
		}
	}

	fmt.Println("✅ Tables created or verified successfully")
}
