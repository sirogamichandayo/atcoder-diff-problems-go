package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	SinDb SinDb
}

type SinDb struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

func LoadConfig() (*Config, error) {
	var config Config
	_, err := toml.DecodeFile("/go/src/app/config/config.toml", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
