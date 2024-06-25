package core

type Datasource interface {
	GetName() string
	GetClient() DatasourceClient
}

const (
	CDatasourceSpotify         = "Spotify"
	CDatasourceYoutube         = "Youtube"
	CTotalSupportedDatasources = 3
)

// helper function usually used when validating user input from network
func IsValidDatasource(datasource string) bool {
	switch datasource {
	// keep up to date with the list of datasources
	case CDatasourceSpotify, CDatasourceYoutube:
		return true
	}
	return false
}
