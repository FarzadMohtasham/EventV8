package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	log.Println("Initializing Database... ⌛")
	database, err := sql.Open("sqlite3", "./api.db")

	if err != nil {
		log.Panic("Could not connect to database")
	}

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)

	DB = database

	createTables()

	log.Println("Database initialized successfully ✅")
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		log.Panic(err)
	}
}
