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
	Mq   MqConfig   `toml:Mq`
}

type HttpConfig struct {
	APIKey string `toml:apikey`
	HOST   string `toml:host`
}

type MqConfig struct {
	HOST string `toml:host`
	PORT string `toml:port`
}
