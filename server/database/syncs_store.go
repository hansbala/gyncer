package database

import (
	"database/sql"
	"errors"
	"strings"
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
		id IN (?)`
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
	if !core.IsValidDatasource(sync.SourceDatasource) {
		return false
	}
	if !core.IsValidDatasource(sync.DestinationDatasource) {
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

// given a list of sync ids returns an array of []Sync
func GetSyncDatas(db *sql.DB, syncIds []string) ([]Sync, error) {
	stmt, err := db.Prepare(cGetSyncDataQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(strings.Join(syncIds, ","))
	if err != nil {
		return nil, err
	}
	syncDatas := make([]Sync, 0)
	for rows.Next() {
		var syncData Sync
		// TODO(Hans): make sure this is okay
		if err := rows.Scan(&syncData); err != nil {
			return nil, err
		}
		syncDatas = append(syncDatas, syncData)
	}
	if len(syncDatas) != len(syncIds) {
		return nil, errors.New("mismatch between number of sync datas and sync ids")
	}
	return syncDatas, nil
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
