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
func Init(app, logLevel string) (*logrus.Entry, error) {
	var err error
	//Default.Default.Formatter = &logrus.TextFormatter{ForceColors: true}
	Default.Logger.Level, err = logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse log level - "+logLevel)
	}

	Default = Default.WithField("app", app)
	return Default, nil
}

func AddSentyHook(dsn string) {
	sentryHook := sentry.NewHook(dsn, logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel)
	Default.Logger.SetLevel(3)
	Default.Logger.AddHook(sentryHook)
}

func AddFilenameHook() {
	filenameHook := filename.NewHook()
	filenameHook.Field = "file"
	Default.Logger.AddHook(filenameHook)
}
