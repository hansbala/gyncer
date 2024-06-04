package database

import (
	"database/sql"

	"golang.org/x/oauth2"
)

const (
	cInsertSpotifyCredentials = `
	INSERT INTO SpotifyKeys
		(user_id, access_token, refresh_token, expiry)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
	access_token = VALUES(access_token),
	refresh_token = VALUES(refresh_token),
	expiry = VALUES(expiry)
	`

	cGetSpotifyCredentialsForUser = `
	SELECT
		(access_token, refresh_token, expiry)
	FROM
		SpotifyKeys
	WHERE user_id = ?
	`
)

type SpotifyCredential struct {
	*oauth2.Token
	UserId string
}

// Saves new credentials to the database. If the user already has credentials, they are updated.
func SaveSpotifyCredentials(db *sql.DB, credential SpotifyCredential) error {
	stmt, err := db.Prepare(cInsertSpotifyCredentials)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		credential.UserId,
		credential.AccessToken,
		credential.RefreshToken,
		credential.Expiry,
	)
	if err != nil {
		return err
	}
	return nil
}

// Gets the credentials for a user from the database.
func GetSpotifyCredentials(db *sql.DB, userId string) (*SpotifyCredential, error) {
	stmt, err := db.Prepare(cGetSpotifyCredentialsForUser)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var credential SpotifyCredential
	err = stmt.QueryRow(userId).Scan(&credential.AccessToken, &credential.RefreshToken, &credential.Expiry)
	if err != nil {
		return nil, err
	}
	credential.UserId = userId
	return &credential, nil
}
