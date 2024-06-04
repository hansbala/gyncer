package database

import (
	"database/sql"

	"golang.org/x/oauth2"
)

const (
	cInsertYoutubeCredential = `
	INSERT INTO GoogleKeys
		(user_id, access_token, refresh_token, expiry)
	VALUES (?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
	access_token = VALUES(access_token),
	refresh_token = VALUES(refresh_token),
	expiry = VALUES(expiry)
	`
	cGetYoutubeCredentialsForUser = `
	SELECT
		(access_token, refresh_token, expiry)
	FROM
		GoogleKeys
	WHERE user_id = ?
	`
)

type YoutubeCredential struct {
	*oauth2.Token
	UserId string
}

// Saves new credentials to the database
func SaveYoutubeCredentials(db *sql.DB, credential YoutubeCredential) error {
	stmt, err := db.Prepare(cInsertYoutubeCredential)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err = stmt.Exec(
		credential.UserId,
		credential.AccessToken,
		credential.RefreshToken,
		credential.Expiry,
	); err != nil {
		return err
	}
	return nil
}

// Gets the credentials for a user from the database
func GetYoutubeCredentials(db *sql.DB, userId string) (YoutubeCredential, error) {
	stmt, err := db.Prepare(cGetYoutubeCredentialsForUser)
	if err != nil {
		return YoutubeCredential{}, err
	}
	defer stmt.Close()

	var credential YoutubeCredential
	err = stmt.QueryRow(userId).Scan(&credential.AccessToken, &credential.RefreshToken, &credential.Expiry)
	if err != nil {
		return YoutubeCredential{}, err
	}
	return credential, nil
}
