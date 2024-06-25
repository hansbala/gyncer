package spotify

import (
	"github.com/hansbala/gyncer/core"
	"github.com/zmb3/spotify/v2"
)

type SpotifyClient struct {
	client *spotify.Client
}

// compile time check
var _ core.DatasourceClient = &SpotifyClient{}

func NewSpotifyClient() core.DatasourceClient {
	return &SpotifyClient{}
}

func (mp *SpotifyClient) FetchPlaylists(userID string) ([]core.Playlist, error) {
	panic("not implemented")
}

func (mp *SpotifyClient) FetchPlaylistItems(playlistID string) ([]core.Record, error) {
	panic("not implemented")
}

func (mp *SpotifyClient) CreatePlaylist(userID string, playlistName string) (core.Playlist, error) {
	panic("not implemented")
}

func (mp *SpotifyClient) AddItemsToPlaylist(playlistID string, items []core.Record) error {
	panic("not implemented")
}

func (mp *SpotifyClient) RemoveItemsFromPlaylist(playlistID string, items []core.Record) error {
	panic("not implemented")
}
