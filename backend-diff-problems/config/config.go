package config

import (
	"github.com/BurntSushi/toml"
	"path"
)

type Config struct {
	SinDb SinDb
	ApiV1 ApiV1
}

type SinDb struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

type ApiV1 struct {
	AllowOrigins []string `toml:"allow-origins"`
}

func LoadConfig(execDir string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path.Join(execDir, "config/config.toml"), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
