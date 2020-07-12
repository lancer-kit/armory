package db

import (
	"os"
	"testing"

	"github.com/lancer-kit/noble"
	"github.com/stretchr/testify/assert"
)

const connURL = "postgres://user:p4s5w0rd@127.0.0.1/exchange?sslmode=disable"

func testDBConfig() SecureConfig {
	return SecureConfig{
		Driver:      "postgres",
		Name:        "exchange",
		Host:        "127.0.0.1",
		User:        noble.Secret{}.New("dynenv:USER_NAME"),
		Password:    noble.Secret{}.New("dynenv:PASSWORD"),
		SSL:         false,
		InitTimeout: 60,
		AutoMigrate: true,
		WaitForDB:   true,
		Params: &ConnectionParams{
			MaxIdleConns: 100,
			MaxOpenConns: 300,
			MaxLifetime:  3200,
		},
	}
}

func TestDBConfig_ConnectionString(t *testing.T) {
	os.Setenv("USER_NAME", "user")
	os.Setenv("PASSWORD", "p4s5w0rd")
	cs := testDBConfig().ConnectionString()
	println(cs)
	assert.Equal(t, connURL, cs)
}

func TestDBConfig_Config(t *testing.T) {
	os.Setenv("USER_NAME", "user")
	os.Setenv("PASSWORD", "p4s5w0rd")
	cfg := testDBConfig().Config()
	assert.NotEmpty(t, cfg)
	assert.Equal(t, connURL, cfg.ConnURL)
	assert.NoError(t, cfg.Validate())
}

func TestDBConfig_Validate(t *testing.T) {
	os.Setenv("USER_NAME", "")
	os.Setenv("PASSWORD", "")
	td := testDBConfig()
	e := td.Validate()
	if !assert.Error(t, e) {
		return
	}
	println(e.Error())
	os.Setenv("USER_NAME", "user")
	os.Setenv("PASSWORD", "p4s5w0rd")
	e = td.Validate()
	assert.NoError(t, e)
	assert.Equal(t, connURL, td.ConnectionString())
}
