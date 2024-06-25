package core

// Record represents a single song.
type Record interface {
	GetId() string
	GetName() string
	GetArtist() string
	GetAlbum() string
}
