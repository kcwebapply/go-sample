package config

import (
	"github.com/BurntSushi/toml"
)

var config Config

func init() {
	GetConfig()
}

func GetConfig() Config {
	toml.DecodeFile("config.toml", &config)
	return config
}

type Config struct {
	Http HttpConfig `toml:Http`
}

type HttpConfig struct {
	APIKey string `toml:apikey`
	HOST   string `toml:host`
}
