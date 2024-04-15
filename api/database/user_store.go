package database

import (
	"database/sql"
	"gyncer/core"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var INSERT_USER_QUERY = `INSERT INTO Users (id, email, hashed_password) VALUES (?, ?, ?)`
var GET_USER_QUERY = `SELECT email, hashed_password FROM Users WHERE id = ? LIMIT 1`

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateNewUserInDB(db *sql.DB, newUser *UserCredentials) error {
	user_hash, err := core.HashString(newUser.Email)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// create the bcrypt hashed password
	// salt is automatically generated and stored in the hash
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// prepare the insert statement
	stmt, err := db.Prepare(INSERT_USER_QUERY)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	// execute the insert statement
	_, err = stmt.Exec(user_hash, newUser.Email, hashed_password)
	if err != nil {
		return err
	}

	return nil
}

func IsValidUserInDB(db *sql.DB, currentUser *UserCredentials) (bool, error) {
	user_hash, err := core.HashString(currentUser.Email)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	// prepare to get the user from db
	stmt, err := db.Prepare(GET_USER_QUERY)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	defer stmt.Close()

	// execute getting the user
	var email, hashed_password string
	if err = stmt.QueryRow(user_hash).Scan(&email, &hashed_password); err != nil {
		if err == sql.ErrNoRows {
			// user does not exist, so return false but no error
			return false, nil
		}
		return false, err
	}

	// compare hashed_password with bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(currentUser.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
