package log

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Default is a log.Entry singleton.
var defaultLog *logrus.Entry // nolint:gochecknoglobals

// nolint:gochecknoinits
func init() {
	l := logrus.New()
	l.Level = logrus.InfoLevel
	host, err := os.Hostname()
	if err != nil {
		logrus.Error(err)
	}
	defaultLog = logrus.NewEntry(l).WithField("hostname", host)
}

// Init initializes a default logger configuration by passed configuration.
func Init(config Config) (*logrus.Entry, error) {
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level - "+config.Level)
	}
	defaultLog.Logger.SetLevel(level)
	defaultLog = defaultLog.WithField("app", config.AppName)

	if config.JSON {
		defaultLog.Logger.Formatter = &logrus.JSONFormatter{}
	}

	if config.Sentry != "" {
		AddSentryHook(config.Sentry)
	}

	return defaultLog, nil
}

// Get is a getter for the `logrus.Entry` singleton.
func Get() *logrus.Entry {
	return defaultLog
}
