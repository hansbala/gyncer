package youtube

import "github.com/hansbala/gyncer/core"

type YoutubeClient struct{}

// compile time check
var _ core.DatasourceClient = &YoutubeClient{}

func (yt *YoutubeClient) FetchPlaylists(userID string) ([]core.Playlist, error) {
	panic("not implemented")
}

func (yt *YoutubeClient) FetchPlaylistItems(playlistID string) ([]core.Record, error) {
	panic("not implemented")
}

func (yt *YoutubeClient) CreatePlaylist(userID string, playlistName string) (core.Playlist, error) {
	panic("not implemented")
}

func (yt *YoutubeClient) AddItemsToPlaylist(playlistID string, items []core.Record) error {
	panic("not implemented")
}

func (yt *YoutubeClient) RemoveItemsFromPlaylist(playlistID string, items []core.Record) error {
	panic("not implemented")
}
