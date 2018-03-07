package log

import (
	"os"

	"github.com/onrik/logrus/filename"
	"github.com/onrik/logrus/sentry"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Default is a log.Entry singleton.
var Default *logrus.Entry

func init() {
	l := logrus.New()
	l.Level = logrus.InfoLevel
	host, _ := os.Hostname()
	Default = logrus.NewEntry(l).WithField("hostname", host)
}

// Init initializes a default logger configuration.
func Init(config Config) (*logrus.Entry, error) {
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level - "+config.Level)
	}
	Default.Logger.SetLevel(level)
	Default = Default.WithField("app", config.AppName)

	if config.AddTrace {
		AddFilenameHook()
	}

	if config.Sentry != "" {
		AddSentyHook(config.Sentry)
	}

	if config.JSON {
		Default.Logger.Formatter = &logrus.JSONFormatter{}
	}

	return Default, nil
}

func AddSentyHook(dsn string) {
	sentryHook := sentry.NewHook(dsn, logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel)
	Default.Logger.AddHook(sentryHook)
}

func AddFilenameHook() {
	filenameHook := filename.NewHook()
	filenameHook.Field = "file"
	Default.Logger.AddHook(filenameHook)
}
