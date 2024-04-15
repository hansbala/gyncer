package database

import (
	"database/sql"
	"errors"
	"gyncer/core"
)

type Sync struct {
	UserID                string `json:"userID"`
	SourceDatasource      string `json:"sourceDatasource"`
	SourcePlaylistId      string `json:"sourcePlaylistId"`
	DestinationDatasource string `json:"destinationDatasource"`
	DestinationPlaylistId string `json:"destinationPlaylistId"`
	SyncFrequency         int32  `json:"syncFrequency"`
}

func (sync *Sync) IsValidSync() bool {
	if !core.Datasource(sync.SourceDatasource).IsValidDatasource() {
		return false
	}
	if !core.Datasource(sync.DestinationDatasource).IsValidDatasource() {
		return false
	}
	return sync.SyncFrequency > 0
	// TODO: validate playlist ids and maybe user id?
}

var INSERT_SYNC_QUERY = `
	INSERT INTO Syncs (
		user_id,
		source_datasource,
		source_playlist_id,
		destination_datasource,
		destination_playlist_id,
		sync_frequency
	) VALUES (?, ?, ?, ?, ?, ?)
`

// insert a new sync into the Syncs table
func InsertNewSync(db *sql.DB, newSync *Sync) error {
	stmt, err := db.Prepare(INSERT_SYNC_QUERY)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		newSync.UserID,
		newSync.SourceDatasource,
		newSync.SourcePlaylistId,
		newSync.DestinationDatasource,
		newSync.DestinationPlaylistId,
		newSync.SyncFrequency,
	)
	if err != nil {
		return err
	}

	numRowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if numRowsAffected != 1 {
		return errors.New("expected to update one row")
	}

	return nil
}
