package core

// represents a type of music streaming service
type Datasource string

const (
	DatasourceSpotify Datasource = "Spotify"
	DatasourceYoutube Datasource = "Youtube"
	DatasourceTidal   Datasource = "Tidal"
)

// helper function usually used when validating user input from network
func (datasource Datasource) IsValidDatasource() bool {
	switch datasource {
	// keep up to date with the list of datasources
	case DatasourceSpotify, DatasourceYoutube, DatasourceTidal:
		return true
	}
	return false
}
