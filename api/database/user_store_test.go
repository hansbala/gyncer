package database_test

import (
	"gyncer/core"
	"gyncer/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUserInDB(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new user
	newUser := database.UserCredentials{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Hash the user's email to simulate what the function does
	user_hash, _ := core.HashString(newUser.Email)

	// Expect a prepare followed by an exec
	mock.ExpectPrepare(regexp.QuoteMeta(database.INSERT_USER_QUERY)).
		ExpectExec().
		WithArgs(user_hash, newUser.Email, sqlmock.AnyArg()). // Use sqlmock.AnyArg() to match any bcrypt hash
		WillReturnResult(sqlmock.NewResult(1, 1))             // Simulate a successful insert with new ID 1 and 1 affected row

	// Call the function to test
	err = database.CreateNewUserInDB(db, &newUser)

	// Make sure there were no errors
	assert.NoError(t, err)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
