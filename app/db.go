package app

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func InitDB(dsn string) *sql.DB {
	// Open the database
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
	createTableUsers(db)
	createTableSession(db)
	createTableSubject(db)
	createTableTag(db)
	createTableQuestion(db)
	createTableResponse(db)
	createTableTagged(db)
	createTablePrecise(db)
	createTableVote_response(db)
	createTableVote_question(db)
}

/* --------- Create Funcs ----------- */
func createTableUsers(db *sql.DB) {
	// Create a User table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS users (
		id_student SERIAL NOT NULL,
		lastname VARCHAR(50) NOT NULL,
		firstname VARCHAR(50) NOT NULL,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		password VARCHAR(250) NOT NULL,
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
		UNIQUE(username),
		UNIQUE(email)
	);`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableSubject(db *sql.DB) {
	// Create a Subject table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Subject(
		id_subject SERIAL NOT NULL,
		title VARCHAR(50) NOT NULL,
		description VARCHAR(500) NOT NULL,
		creation_date DATE NOT NULL,
		update_date DATE NOT NULL,
		PRIMARY KEY(id_subject),
		UNIQUE(title)
	);`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableTag(db *sql.DB) {
	// Create a Tag table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Tag(
		id_tag SERIAL NOT NULL,
		title VARCHAR(50) NOT NULL,
		description VARCHAR(500) NOT NULL,
		creation_date DATE NOT NULL,
		update_date DATE NOT NULL,
		PRIMARY KEY(id_tag),
		UNIQUE(title)
	);`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableQuestion(db *sql.DB) {
	// Create a Question table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Question(
		id_question SERIAL NOT NULL,
		title VARCHAR(50) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		content VARCHAR(10000) NOT NULL,
		upvotes INT,
		downvotes INT,
		creation_date DATE NOT NULL,
		update_date DATE NOT NULL,
		id_student INT NOT NULL,
		id_subject INT NOT NULL,
		PRIMARY KEY(id_question),
		UNIQUE(title),
		FOREIGN KEY(id_student) REFERENCES users(id_student),
		FOREIGN KEY(id_subject) REFERENCES Subject(id_subject)
	);`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableResponse(db *sql.DB) {
	// Create a Response table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Response(
		id_response SERIAL NOT NULL,
		description VARCHAR(1000) NOT NULL,
		content VARCHAR(5000) NOT NULL,
		upvotes INT,
		downvotes INT,
		best_answer BOOLEAN NOT NULL,
		creation_date DATE NOT NULL,
		update_date DATE NOT NULL,
		id_question INT NOT NULL,
		id_student INT NOT NULL,
		PRIMARY KEY(id_response),
		FOREIGN KEY(id_question) REFERENCES Question(id_question) ON DELETE CASCADE,
		FOREIGN KEY(id_student) REFERENCES users(id_student)
	);
	`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableTagged(db *sql.DB) {
	// Create a Tagged table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Tagged(
		id_question INT,
		id_tag INT,
		PRIMARY KEY(id_question, id_tag),
		FOREIGN KEY(id_question) REFERENCES Question(id_question),
		FOREIGN KEY(id_tag) REFERENCES Tag(id_tag)
	);
	`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTablePrecise(db *sql.DB) {
	// Create a Precise table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Precise(
		id_subject INT,
		id_tag INT,
		PRIMARY KEY(id_subject, id_tag),
		FOREIGN KEY(id_subject) REFERENCES Subject(id_subject),
		FOREIGN KEY(id_tag) REFERENCES Tag(id_tag)
	);
	`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableVote_response(db *sql.DB) {
	// Create a	Vote_response table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Vote_response(
		id_student INT,
		id_response INT,
		upvote_r BOOLEAN NOT NULL,
		downvote_r BOOLEAN NOT NULL,
		PRIMARY KEY(id_student, id_response),
		FOREIGN KEY(id_student) REFERENCES users(id_student),
		FOREIGN KEY(id_response) REFERENCES Response(id_response)
	);
	`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableVote_question(db *sql.DB) {
	// Create a	Vote_question table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Vote_question(
		id_student INT,
		id_question INT,
		upvote_q BOOLEAN NOT NULL,
		downvote_q BOOLEAN NOT NULL,
		PRIMARY KEY(id_student, id_question),
		FOREIGN KEY(id_student) REFERENCES users(id_student),
		FOREIGN KEY(id_question) REFERENCES Question(id_question) ON DELETE CASCADE
	);
	`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func createTableSession(db *sql.DB) {
	// Create a	Vote_question table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Sessions(
		id SERIAL NOT NULL,
		uuid VARCHAR(50) NOT NULL,
		user_id INT NOT NULL,
		expire_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL
	);
	`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

// insertSessionToDB inserts a new session into the database.
// It takes a database connection, user ID, user UUID, creation timestamp, and expiration timestamp as parameters.
// It returns an error if there was a problem inserting the session.
func insertSessionToDB(db *sql.DB, user_id int, user_uuid string, createdAt time.Time, expireAt time.Time) error {
	stmt, err := db.Prepare("INSERT INTO Sessions(user_id,uuid,created_at,expire_at) VALUES($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(user_id, user_uuid, createdAt, expireAt); err != nil {
		return err
	}
	return err
}

// getUserIDFromDB retrieves the user ID from the database based on the given username.
// It takes the username and a pointer to the SQL database connection as input parameters.
// It returns the user ID as an integer and an error if any occurred during the database operation.
func getUserIDFromDB(username string, db *sql.DB) (int, error) {
	var id int
	stmt, err := db.Prepare("SELECT id_student FROM users WHERE username = $1 OR email = $2")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username, username).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return 0, err
	case err != nil:
		return 0, err
	}
	return id, nil
}

func getUserIDUsingSessionID(sessionID string, db *sql.DB) (int, error) {
	var id int
	stmt, err := db.Prepare("SELECT user_id FROM Sessions WHERE uuid = $1")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(sessionID).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		return 0, err
	case err != nil:
		return 0, err
	}
	return id, nil
}
