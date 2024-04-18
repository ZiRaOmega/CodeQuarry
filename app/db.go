package app

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func InitDB(filepath string) *sql.DB {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database is accessible by pinging it
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func SetupDB(db *sql.DB) {
	// Create a User table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS users (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"fullname" TEXT,
		"username" TEXT UNIQUE,
		"email" TEXT UNIQUE,
		"password" TEXT
	);`

	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
