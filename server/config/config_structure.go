package config

type Config struct {
	Database Database `toml:"database"`
	Server   Server   `toml:"server"`
	Spotify  Spotify  `toml:"spotify"`
}

type Database struct {
	MySqlRootUser     string `toml:"mysql_root_user"`
	MySqlRootPassword string `toml:"mysql_root_password"`
	MySqlUser         string `toml:"mysql_user"`
	MySqlPassword     string `toml:"mysql_password"`
	MySqlDatabase     string `toml:"mysql_database"`
	MySqlPort         int    `toml:"mysql_port"`
}

type Server struct {
	JwtSecret         string `toml:"jwt_secret"`
	GinMode           string `toml:"gin_mode"`
	MetaSyncFrequency int    `toml:"meta_sync_frequency"`
}

type Spotify struct {
	GyncerClientId     string `toml:"gyncer_client_id"`
	GyncerClientSecret string `toml:"gyncer_client_secret"`
	GyncerRedirectUrl  string `toml:"gyncer_redirect_url"`
}
