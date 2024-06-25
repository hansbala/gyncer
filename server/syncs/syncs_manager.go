package syncs

import (
	"github.com/hansbala/gyncer/core"
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
	// create clients for the source and destination
	sourceDatasource, err := core.GetDatasourceFactoryRegistry().GetDatasource(sync.SourceDatasource)
	if err != nil {
		return err
	}
	sourceClient := sourceDatasource.GetClient()
	destinationDatasource, err := core.GetDatasourceFactoryRegistry().GetDatasource(sync.DestinationDatasource)
	if err != nil {
		return err
	}
	destinationClient := destinationDatasource.GetClient()
	// TODO: We should also include some field to define the type of merge taking place:
	// 1. ONE-WAY MERGE -> Source goes to destination. Anything extra in destination is deleted
	// 2. TWO-WAY MERGE -> Source + destination goes to both (superset of both)
	// ... Probably some more options but these suit my use-cases.
	sourceClient.FetchPlaylistItems(sync.SourcePlaylistId)
	destinationClient.FetchPlaylistItems(sync.DestinationPlaylistId)

	return nil
}
