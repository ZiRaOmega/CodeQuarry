package app

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func InitDB(dsn string) *sql.DB {
	// Open the SQLite database
	db, err := sql.Open("postgres", dsn)
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
	id_student SERIAL NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	firstname VARCHAR(50) NOT NULL,
	nickname VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password VARCHAR(50) NOT NULL,
	avatar VARCHAR(50),
	birth_date DATE,
	bio VARCHAR(100),
	website VARCHAR(50),
	github VARCHAR(50),
	xp INT,
	rang_rank_ INT,
	school_year DATE,
	creation_date DATE,
	update_date DATE,
	deleting_date DATE,
	PRIMARY KEY(id_student),
	UNIQUE(nickname),
	UNIQUE(email)
);`

	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
