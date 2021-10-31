package config

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	LoadConfigs()
}

func TestLoadAppConfig(t *testing.T) {
	assert.NotEmpty(t, strconv.FormatBool(App.Debug))
	assert.NotEmpty(t, App.Env)
	assert.NotEmpty(t, App.Port)
}

func TestLoadDbConfig(t *testing.T) {
	assert.NotEmpty(t, Db.Host)
	assert.NotEmpty(t, Db.Port)
	assert.NotEmpty(t, Db.Database)
	assert.NotEmpty(t, Db.Username)
	assert.NotNil(t, Db.Password)
}
