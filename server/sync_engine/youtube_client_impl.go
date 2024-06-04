package sync_engine

type YoutubeClient struct{}

// compile time check
var _ MusicProvider = &YoutubeClient{}

func (mp *YoutubeClient) FetchPlaylists(userID string) ([]Playlist, error) {
	panic("not implemented")
}

func (mp *YoutubeClient) FetchPlaylistItems(playlistID string) ([]PlaylistItem, error) {
	panic("not implemented")
}

func (mp *YoutubeClient) CreatePlaylist(userID string, playlistName string) (Playlist, error) {
	panic("not implemented")
}

func (mp *YoutubeClient) AddItemsToPlaylist(playlistID string, items []PlaylistItem) error {
	panic("not implemented")
}

func (mp *YoutubeClient) RemoveItemsFromPlaylist(playlistID string, items []PlaylistItem) error {
	panic("not implemented")
}
