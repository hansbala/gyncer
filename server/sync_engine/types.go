package sync_engine

type MusicProvider interface {
	FetchPlaylists(userID string) ([]Playlist, error)
	FetchPlaylistItems(playlistID string) ([]PlaylistItem, error)
	CreatePlaylist(userID string, playlistName string) (Playlist, error)
	AddItemsToPlaylist(playlistID string, items []PlaylistItem) error
	RemoveItemsFromPlaylist(playlistID string, items []PlaylistItem) error
}

// Playlist represents a generic playlist structure.
type Playlist struct {
	ID    string
	Name  string
	Items []PlaylistItem
}

// PlaylistItem represents a generic playlist item structure.
type PlaylistItem struct {
	ID     string
	Name   string
	Artist string
	Album  string
}
