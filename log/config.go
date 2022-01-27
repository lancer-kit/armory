package log

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lancer-kit/noble"
)

// Config is an options for the initialization
// of the default logrus.Entry.
type Config struct {
	// AppName identifier of the app.
	AppName string `json:"app_name" yaml:"app_name"`
	// Level is a string representation of the `lorgus.Level`.
	Level string `json:"level" yaml:"level"`
	// Sentry is a DSN string for sentry hook.
	Sentry string `json:"sentry" yaml:"sentry"`
	// AddTrace enable adding of the filename field into log.
	// DEPRECATED
	AddTrace bool `json:"add_trace" yaml:"add_trace"`
	// JSON enable json formatted output.
	JSON bool `json:"json" yaml:"json"`
}

// NConfig is an options for the initialization
// of the default logrus.Entry.
// N stands for Noble.
type NConfig struct {
	// AppName identifier of the app.
	AppName string `yaml:"app_name"`
	// Level is a string representation of the `lorgus.Level`.
	Level noble.Secret `yaml:"level"`
	// Sentry is a DSN string for sentry hook.
	Sentry noble.Secret `yaml:"sentry"`
	// AddTrace enable adding of the filename field into log.
	AddTrace bool `yaml:"add_trace"`
	// JSON enable json formatted output.
	JSON bool `yaml:"json"`
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (cfg NConfig) Validate() error {
	return validation.ValidateStruct(&cfg,
		validation.Field(&cfg.AppName, validation.Required),
		validation.Field(&cfg.Level, validation.Required, noble.RequiredSecret),
	)
}
