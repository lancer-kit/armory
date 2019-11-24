package db

import validation "github.com/go-ozzo/ozzo-validation"

type ConnectionParams struct {
	MaxIdleConns int `json:"max_idle" yaml:"max_idle"`
	MaxOpenConns int `json:"max_open" yaml:"max_open"`
	// MaxLifetime time.Duration in Millisecond
	MaxLifetime int64 `json:"max_lifetime" yaml:"max_lifetime"`
}

type Config struct {
	ConnURL     string `json:"conn_url" yaml:"conn_url"` //The database connection string.
	InitTimeout int    `json:"dbInitTimeout" yaml:"init_timeout"`
	// AutoMigrate if `true` execute db migrate up on start.
	AutoMigrate bool              `json:"auto_migrate" yaml:"auto_migrate"`
	WaitForDB   bool              `json:"wait_for_db" yaml:"wait_for_db"`
	Params      *ConnectionParams `json:"params" yaml:"params"`
}

func (cfg *Config) URL() string {
	return cfg.ConnURL
}

func (cfg Config) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.ConnURL, validation.Required),
		validation.Field(&cfg.InitTimeout, validation.Required),
	)
}
