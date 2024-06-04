package sync_engine

import "github.com/zmb3/spotify/v2"

type SpotifyClient struct {
	client *spotify.Client
}

// compile time check
var _ MusicProvider = &SpotifyClient{}

func (mp *SpotifyClient) FetchPlaylists(userID string) ([]Playlist, error) {
	panic("not implemented")
}

func (mp *SpotifyClient) FetchPlaylistItems(playlistID string) ([]PlaylistItem, error) {
	panic("not implemented")
}

func (mp *SpotifyClient) CreatePlaylist(userID string, playlistName string) (Playlist, error) {
	panic("not implemented")
}

func (mp *SpotifyClient) AddItemsToPlaylist(playlistID string, items []PlaylistItem) error {
	panic("not implemented")
}

func (mp *SpotifyClient) RemoveItemsFromPlaylist(playlistID string, items []PlaylistItem) error {
	panic("not implemented")
}
