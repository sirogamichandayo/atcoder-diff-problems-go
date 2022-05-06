package config

import (
	"os"
	"strings"
)

type Config struct {
	SinDb SinDb
	ApiV1 ApiV1
}

type SinDb struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type ApiV1 struct {
	AllowOrigin string
}

func LoadConfig(execDir string) (*Config, error) {
	envStrList := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		envStrList[pair[0]] = pair[1]
	}

	return &Config{
		SinDb: loadSinDb(envStrList),
		ApiV1: loadApiV1(envStrList),
	}, nil
}

func loadSinDb(envStrList map[string]string) SinDb {
	return SinDb{
		Host:     envStrList["SIN_DB_HOST"],
		Port:     envStrList["SIN_DB_PORT"],
		User:     envStrList["SIN_DB_USER"],
		Password: envStrList["SIN_DB_PASSWORD"],
		Database: envStrList["SIN_DB_DATABASE"],
	}
}

func loadApiV1(envStrList map[string]string) ApiV1 {
	return ApiV1{
		AllowOrigin: envStrList["API_ALLOW_ORIGIN"],
	}
}
