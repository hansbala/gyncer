package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/hansbala/gyncer/core"
)

const (
	cInsertSyncQuery = `
	INSERT INTO Syncs (
		user_id,
		source_datasource,
		source_playlist_id,
		destination_datasource,
		destination_playlist_id,
		sync_frequency
	) VALUES (?, ?, ?, ?, ?, ?)
`

	cGetSyncsToSyncQuery = `
	SELECT
		id
	FROM
		Syncs 
	WHERE
		last_synced_at = NULL OR DATE_ADD(last_synced_at, INTERVAL sync_frequency HOUR) < ?`

	cGetSyncDataQuery = `
	SELECT
		id,
		user_id, 
		source_datasource,
		source_playlist_id,
		destination_datasource,
		destination_playlist_id,
		sync_frequency 
	FROM
		Syncs
	WHERE
		id = ?
	LIMIT
		1`
)

type Sync struct {
	UserID                string `json:"userID"`
	SourceDatasource      string `json:"sourceDatasource"`
	SourcePlaylistId      string `json:"sourcePlaylistId"`
	DestinationDatasource string `json:"destinationDatasource"`
	DestinationPlaylistId string `json:"destinationPlaylistId"`
	SyncFrequency         int32  `json:"syncFrequency"`
}

type StartSync struct {
	SyncIds []string `json:"syncIds"`
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

// insert a new sync into the Syncs table
func InsertNewSync(db *sql.DB, newSync *Sync) error {
	stmt, err := db.Prepare(cInsertSyncQuery)
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

func GetSyncData(db *sql.DB, syncId string) (*Sync, error) {
	stmt, err := db.Prepare(cGetSyncDataQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(syncId)
	if err != nil {
		return nil, err
	}
	var syncData Sync
	for rows.Next() {
		// TODO(Hans): Serialize data before reading
		if err := rows.Scan(&syncData); err != nil {
			return nil, err
		}
		// TODO(Hans): static analysis
		return &syncData, nil
	}
	return nil, errors.New("failed to get sync data from SQL")
}

// based on the time provided, returns the sync ids that need to be synced
func GetSyncsToSync(db *sql.DB, currentTime time.Time) ([]int, error) {
	stmt, err := db.Prepare(cGetSyncsToSyncQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Query(currentTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}
	syncIds := make([]int, 0)
	for res.Next() {
		var syncId int
		err := res.Scan(&syncId)
		if err != nil {
			return nil, err
		}
		syncIds = append(syncIds, syncId)
	}
	return syncIds, nil
}
