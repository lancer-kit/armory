package db

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lancer-kit/noble"
)

const (
	DriverPostgres = "postgres"
	DriverMySQL    = "mysql"
)

type Config struct {
	ConnURL     string            `json:"conn_url" yaml:"conn_url"` // The database connection string.
	InitTimeout int               `json:"dbInitTimeout" yaml:"init_timeout"`
	AutoMigrate bool              `json:"auto_migrate" yaml:"auto_migrate"`
	WaitForDB   bool              `json:"wait_for_db" yaml:"wait_for_db"`
	Params      *ConnectionParams `json:"params" yaml:"params"`
}

func (cfg *Config) URL() string {
	return cfg.ConnURL
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (cfg Config) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.ConnURL, validation.Required),
		validation.Field(&cfg.InitTimeout, validation.Required),
	)
}

type ConnectionParams struct {
	MaxIdleConns int `json:"max_idle" yaml:"max_idle"`
	MaxOpenConns int `json:"max_open" yaml:"max_open"`
	// MaxLifetime time.Duration in Millisecond
	MaxLifetime int64 `json:"max_lifetime" yaml:"max_lifetime"`
}

// SecureConfig configuration with secrets support
// nolint:maligned
type SecureConfig struct {
	Driver      string            `yaml:"driver" json:"driver"`
	Name        string            `yaml:"name" json:"name"`
	Host        string            `yaml:"host" json:"host"`
	Port        int               `yaml:"port"  json:"port"`
	User        noble.Secret      `yaml:"user" json:"user"`
	Password    noble.Secret      `yaml:"password" json:"password"`
	SSL         bool              `yaml:"ssl" json:"ssl"`
	MySQLParams string            `yaml:"my_sql_params" json:"my_sql_params,omitempty"`
	InitTimeout int               `yaml:"init_timeout" json:"init_timeout"`
	AutoMigrate bool              `yaml:"auto_migrate" json:"auto_migrate"`
	WaitForDB   bool              `yaml:"wait_for_db" json:"wait_for_db"`
	Params      *ConnectionParams `yaml:"params" json:"params"`
}

// ConnectionString returns Connection String for selected driver
func (d SecureConfig) ConnectionString() string {
	port := ""
	if d.Port != 0 {
		port = fmt.Sprintf(":%d", d.Port)
	}

	switch d.Driver {
	case DriverPostgres:
		mode := ""
		if !d.SSL {
			mode = "?sslmode=disable"
		}
		DSN := `postgres://%s:%s@%s%s/%s%s`
		return fmt.Sprintf(DSN, d.User.Get(), d.Password.Get(), d.Host, port, d.Name, mode)
	case DriverMySQL:
		// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
		params := ""
		if d.MySQLParams != "" {
			params = "?" + d.MySQLParams
		}
		DSN := `%s:%s@tcp(%s%s)/%s%s`
		return fmt.Sprintf(DSN, d.User.Get(), d.Password.Get(), d.Host, port, d.Name, params)
	}
	return ""
}

// Config returns lancer db Config
func (d SecureConfig) Config() Config {
	return Config{
		ConnURL:     d.ConnectionString(),
		InitTimeout: d.InitTimeout,
		AutoMigrate: d.AutoMigrate,
		WaitForDB:   d.WaitForDB,
		Params:      d.Params,
	}
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (d SecureConfig) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Driver, validation.Required),
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Host, validation.Required),
		validation.Field(&d.InitTimeout, validation.Required),
		validation.Field(&d.User, validation.Required, noble.RequiredSecret),
		validation.Field(&d.Password, validation.Required, noble.RequiredSecret),
	)
}
