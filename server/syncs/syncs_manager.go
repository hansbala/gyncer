package syncs

import (
	"github.com/hansbala/gyncer/database"
)

// receives a list of syncs to perform, updates the database for each one.
// only returns error, void if all okay
func StartSyncWrapper(syncDatas []database.Sync) error {
	for _, sync := range syncDatas {
		if err := startSync(sync); err != nil {
			return err
		}
	}
	return nil
}

// receives one sync and starts it
// TODO(Hans): This can be it's own go routine so no need to block execution of main thread
func startSync(sync database.Sync) error {
	// TODO: Handle vendor syncing here
	panic("not implemented")
}
