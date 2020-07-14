package log

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Entry is a type alias for `*logrus.Entry`,
// can be used to avoid direct import of the `logrus` package.
// DEPRECATED
type Entry = *logrus.Entry

// Default is a log.Entry singleton.
// DEPRECATED
var Default *logrus.Entry // nolint:gochecknoglobals

// nolint:gochecknoinits
func init() {
	l := logrus.New()
	l.Level = logrus.InfoLevel
	host, err := os.Hostname()
	if err != nil {
		logrus.Error(err)
	}
	Default = logrus.NewEntry(l).WithField("hostname", host)
}

// Init initializes a default logger configuration by passed configuration.
func Init(config Config) (*logrus.Entry, error) {
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level - "+config.Level)
	}
	Default.Logger.SetLevel(level)
	Default = Default.WithField("app", config.AppName)

	if config.JSON {
		Default.Logger.Formatter = &logrus.JSONFormatter{}
	}

	if config.AddTrace {
		AddFilenameHook()
	}

	if config.Sentry != "" {
		AddSentryHook(config.Sentry)
	}

	return Default, nil
}

// Get is a getter for the `logrus.Entry` singleton.
func Get() *logrus.Entry {
	return Default
}
