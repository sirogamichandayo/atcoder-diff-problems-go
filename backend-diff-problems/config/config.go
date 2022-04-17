package config

import (
	"github.com/BurntSushi/toml"
	"path"
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

func LoadConfig(execDir string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path.Join(execDir, "config/config.toml"), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
