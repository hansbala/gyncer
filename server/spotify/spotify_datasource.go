package spotify

import "github.com/hansbala/gyncer/core"

type spotifyDatasourceImpl struct{}

var _ core.Datasource = &spotifyDatasourceImpl{}

func init() {
	core.GetDatasourceFactoryRegistry().RegisterDatasource(core.CDatasourceSpotify, func() core.Datasource {
		return &spotifyDatasourceImpl{}
	})
}

func (sd *spotifyDatasourceImpl) GetName() string {
	return core.CDatasourceSpotify
}
func (sd *spotifyDatasourceImpl) GetClient() core.DatasourceClient {
	return NewSpotifyClient()
}
