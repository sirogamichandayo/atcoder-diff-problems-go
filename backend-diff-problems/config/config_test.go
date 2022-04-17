package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var config *Config

func TestConfigMain(t *testing.T) {
	var err error
	config, err = LoadConfig()
	assert.Nil(t, err)

	t.Run("SinDbの値の存在チェック", SinDbの値の存在チェック)
}

func SinDbの値の存在チェック(t *testing.T) {
	sinDb := config.SinDb
	assert.NotEmpty(t, sinDb.Host)
	assert.NotEmpty(t, sinDb.Port)
	assert.NotEmpty(t, sinDb.User)
	assert.NotEmpty(t, sinDb.Password)
	assert.NotEmpty(t, sinDb.Database)
}
