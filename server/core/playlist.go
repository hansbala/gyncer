package core

// Playlist represents a generic playlist structure.
type Playlist interface {
	GetId() string
	GetName() string
	GetItems() []Record
}
