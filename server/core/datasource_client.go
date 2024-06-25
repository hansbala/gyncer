package core

type DatasourceClient interface {
	FetchPlaylists(userID string) ([]Playlist, error)
	FetchPlaylistItems(playlistID string) ([]Record, error)
	CreatePlaylist(userID string, playlistName string) (Playlist, error)
	AddItemsToPlaylist(playlistID string, items []Record) error
	RemoveItemsFromPlaylist(playlistID string, items []Record) error
}
