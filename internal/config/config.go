package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	createTables()
	fmt.Println("✅ Database connected successfully")
}

func createTables() {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS events (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			date VARCHAR(255),
			time VARCHAR(255),
			venue VARCHAR(255),
			description TEXT
		);`,

		`CREATE TABLE IF NOT EXISTS attendees (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255),
			event_id INT,
			FOREIGN KEY (event_id) REFERENCES events(id)
		);`,

		`CREATE TABLE IF NOT EXISTS feedback (
			id INT AUTO_INCREMENT PRIMARY KEY,
			event_id INT,
			name VARCHAR(255),
			comment TEXT,
			rating INT,
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
