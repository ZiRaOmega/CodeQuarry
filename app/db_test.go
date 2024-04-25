package app

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInitDB(t *testing.T) {
	// Create a new mock SQL connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Expectations: Expect Ping to be called and to succeed
	mock.ExpectPing()

	// Call InitDB with a mock DSN
	resultDB := InitDB("postgres://user:pass@localhost/dbname")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	// Check that the returned DB is of
	// the correct type
	if resultDB == nil {
		t.Errorf("expected DB connection, got nil")
	}

}

func TestSetupDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Mock expectations for each table creation
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").WillReturnResult(sqlmock.NewResult(1, 1))
	// Add similar expectations for other tables...

	// Call SetupDB
	SetupDB(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateTableUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").WillReturnResult(sqlmock.NewResult(1, 1))

	// Call createTableUsers
	createTableUsers(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Add similar tests for other table creation functions
