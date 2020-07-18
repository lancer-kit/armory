package natsx

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Config is configuration for the interaction with the NATS server.
type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Validate is an implementation of Validatable interface from ozzo-validation.
func (config Config) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.Host, validation.Required),
		validation.Field(&config.Port, validation.Required),
	)
}

// ToURL formats config into NATS connection string.
func (config Config) ToURL() string {
	return fmt.Sprintf("nats://%s:%d", config.Host, config.Port)
}
