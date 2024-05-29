package config

import "github.com/BurntSushi/toml"

var config *Config

func loadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func GetConfig() *Config {
	// if config was already loaded, return it
	if config != nil {
		return config
	}
	// always fetch config file from root of project repo
	configPath := "./config/config.toml"
	receivedConfig, err := loadConfig(configPath)
	if err != nil {
		panic("error loading config file: " + err.Error())
	}
	config = receivedConfig
	return config
}
