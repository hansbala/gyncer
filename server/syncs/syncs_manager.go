package syncs

import (
	"errors"

	"github.com/hansbala/gyncer/core"
	"github.com/hansbala/gyncer/database"
	"github.com/hansbala/gyncer/sync_engine"
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

func getClientFromDatasource(datasource core.Datasource) (sync_engine.MusicProvider, error) {
	switch datasource {
	case core.DatasourceSpotify:
		return &sync_engine.SpotifyClient{}, nil
	case core.DatasourceYoutube:
		return &sync_engine.YoutubeClient{}, nil
	default:
		return nil, errors.New("unknown datasource")
	}

}

// receives one sync and starts it
// TODO(Hans): This can be it's own go routine so no need to block execution of main thread
func startSync(sync database.Sync) error {
	// create clients for the source and destination
	sourceMusicProvider, err := getClientFromDatasource(core.Datasource(sync.SourceDatasource))
	if err != nil {
		return err
	}
	destinationMusicProvider, err := getClientFromDatasource(core.Datasource(sync.DestinationDatasource))
	if err != nil {
		return err
	}
	// TODO: We should also include some field to define the type of merge taking place:
	// 1. ONE-WAY MERGE -> Source goes to destination. Anything extra in destination is deleted
	// 2. TWO-WAY MERGE -> Source + destination goes to both (superset of both)
	// ... Probably some more options but these suit my use-cases.
	sourceMusicProvider.FetchPlaylistItems(sync.SourcePlaylistId)
	destinationMusicProvider.FetchPlaylistItems(sync.DestinationPlaylistId)

	return nil
}
