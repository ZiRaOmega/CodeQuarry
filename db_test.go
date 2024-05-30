package main_test

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"codequarry/app"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// initTestDB initializes a test database
func initTestDB(t *testing.T) *sql.DB {
	// Load environment variables
	err := godotenv.Load()
	require.NoError(t, err, "Error loading .env file")

	// Construct the DSN
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"codequarry",
		"CQ1234",
		"localhost",
		"5432",
		"codequarrytest",
	)

	// Open the database
	db, err := sql.Open("postgres", dsn)
	require.NoError(t, err, "Error opening database")

	// Ping the database to check connectivity
	err = db.Ping()
	require.NoError(t, err, "Error pinging database")

	return db
}

// TestInitDB tests the InitDB function
func TestInitDB(t *testing.T) {
	db := initTestDB(t)
	defer db.Close()

	require.NoError(t, db.Ping())
}

// TestSetupDB tests the SetupDB function
func TestSetupDB(t *testing.T) {
	db := initTestDB(t)
	defer db.Close()

	ResetDB(db)

	tables := []string{
		"users",
		"subject",
		"question",
		"response",
		"favori",
		"vote_question",
		"vote_response",
		"verifyemail",
		"sessions",
	}

	for _, table := range tables {
		var exists bool
		query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s')", strings.ToLower(table))
		err := db.QueryRow(query).Scan(&exists)
		if err != nil {
			fmt.Printf("Error checking table %s: %v\n", table, err)
		}
		require.NoError(t, err, fmt.Sprintf("Error checking table %s", table))
		assert.True(t, exists, fmt.Sprintf("Table %s does not exist", table))
	}
}

func TestGetUserIDFromDB(t *testing.T) {
	db := initTestDB(t)
	defer db.Close()

	username := "test"
	email := "testuser@example.com"
	password := "test"

	_, err := db.Exec("INSERT INTO users (lastname, firstname, username, email, password) VALUES ($1, $2, $3, $4, $5)", "Last", "First", username, email, password)
	require.NoError(t, err)

	id, err := app.GetUserIDFromDB(username, db)
	require.NoError(t, err)
	assert.NotEqual(t, 0, id)

	id, err = app.GetUserIDFromDB(email, db)
	require.NoError(t, err)
	assert.NotEqual(t, 0, id)
}

// TestInsertSessionToDB tests the insertSessionToDB function
func TestInsertSessionToDB(t *testing.T) {
	db := initTestDB(t)
	defer db.Close()

	userID := 1
	userUUID := "test-uuid"
	createdAt := time.Now()
	expireAt := createdAt.Add(24 * time.Hour)

	err := app.InsertSessionToDB(db, userID, userUUID, createdAt, expireAt)
	require.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Sessions WHERE user_id = $1 AND uuid = $2", userID, userUUID).Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

// TestGetUserIDFromDB tests the getUserIDFromDB function

// TestGetUserIDUsingSessionID tests the GetUserIDUsingSessionID function
func TestGetUserIDUsingSessionID(t *testing.T) {
	db := initTestDB(t)
	defer db.Close()
	ResetSession(db)
	userID := 1
	userUUID := "test-uuid"
	createdAt := time.Now()
	expireAt := createdAt.Add(24 * time.Hour)

	_, err := db.Exec("INSERT INTO Sessions (user_id, uuid, created_at, expire_at) VALUES ($1, $2, $3, $4)", userID, userUUID, createdAt, expireAt)
	require.NoError(t, err)

	id, err := app.GetUserIDUsingSessionID(userUUID, db)
	require.NoError(t, err)
	assert.Equal(t, userID, id)
}

// ResetDB resets the database by dropping all tables and recreating them
func ResetDB(db *sql.DB) {
	dropTables(db)
	app.SetupDB(db)
}

// dropTables drops all tables in the database
func dropTables(db *sql.DB) {
	tables := []string{
		"Favori",
		"Vote_response",
		"Vote_question",
		"Response",
		"Question",
		"VerifyEmail",
		"Sessions",
		"Subject",
		"users",
	}

	for _, table := range tables {
		query := "DROP TABLE IF EXISTS " + table + " CASCADE"
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error dropping table %s: %v", table, err)
		}
	}
}

func ResetSession(db *sql.DB) {
	dropSession(db)
	createTableSession(db)
}

func dropSession(db *sql.DB) {
	query := "DROP TABLE IF EXISTS Sessions CASCADE"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Error dropping table Sessions: %v", err)
	}
}

func createTableSession(db *sql.DB) {
	// Create a	Vote_question table
	tableCreationQuery := `CREATE TABLE IF NOT EXISTS Sessions(
		id SERIAL NOT NULL,
		uuid VARCHAR(50) NOT NULL,
		user_id INT NOT NULL,
		expire_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL,
		UNIQUE(uuid),
		PRIMARY KEY(id),
		FOREIGN KEY(user_id) REFERENCES users(id_student) ON DELETE CASCADE
	);
	`
	// Execute the table creation query
	if _, err := db.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
