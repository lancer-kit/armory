package log

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

var (
	levelsMap = map[logrus.Level]sentry.Level{
		logrus.PanicLevel: sentry.LevelFatal,
		logrus.FatalLevel: sentry.LevelFatal,
		logrus.ErrorLevel: sentry.LevelError,
		logrus.WarnLevel:  sentry.LevelWarning,
		logrus.InfoLevel:  sentry.LevelInfo,
		logrus.DebugLevel: sentry.LevelDebug,
		logrus.TraceLevel: sentry.LevelDebug,
	}
)

type SentryOptions sentry.ClientOptions

type SentryHook struct {
	client       *sentry.Client
	levels       []logrus.Level
	tags         map[string]string
	release      string
	environment  string
	prefix       string
	flushTimeout time.Duration
}

func (hook *SentryHook) Levels() []logrus.Level {
	return hook.levels
}

func (hook *SentryHook) Fire(entry *logrus.Entry) error {
	var exceptions []sentry.Exception

	if err, ok := entry.Data[logrus.ErrorKey].(error); ok && err != nil {
		stacktrace := sentry.ExtractStacktrace(err)
		if stacktrace == nil {
			stacktrace = sentry.NewStacktrace()
		}
		exceptions = append(exceptions, sentry.Exception{
			Type:       entry.Message,
			Value:      err.Error(),
			Stacktrace: stacktrace,
		})
	}

	event := sentry.Event{
		Level:       levelsMap[entry.Level],
		Message:     hook.prefix + entry.Message,
		Extra:       map[string]interface{}(entry.Data),
		Tags:        hook.tags,
		Environment: hook.environment,
		Release:     hook.release,
		Exception:   exceptions,
	}

	hub := sentry.CurrentHub()
	hook.client.CaptureEvent(&event, nil, hub.Scope())

	if entry.Level == logrus.PanicLevel || entry.Level == logrus.FatalLevel || entry.Level == logrus.ErrorLevel {
		hook.Flush()
	}

	return nil
}

func (hook *SentryHook) SetPrefix(prefix string) {
	hook.prefix = prefix
}

func (hook *SentryHook) SetTags(tags map[string]string) {
	hook.tags = tags
}

func (hook *SentryHook) AddTag(key, value string) {
	hook.tags[key] = value
}

func (hook *SentryHook) SetRelease(release string) {
	hook.release = release
}

func (hook *SentryHook) SetEnvironment(environment string) {
	hook.environment = environment
}

func (hook *SentryHook) SetFlushTimeout(timeout time.Duration) {
	hook.flushTimeout = timeout
}

func (hook *SentryHook) Flush() {
	hook.client.Flush(hook.flushTimeout)
}

func NewSentryHook(options SentryOptions, levels ...logrus.Level) (*SentryHook, error) {
	client, err := sentry.NewClient(sentry.ClientOptions(options))
	if err != nil {
		return nil, err
	}

	hook := SentryHook{
		client:       client,
		levels:       levels,
		tags:         map[string]string{},
		flushTimeout: 10 * time.Second,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}

	return &hook, nil
}
