package log

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// AddSentryHook adds hook that sends all error,
// fatal and panic log lines to the sentry service.
func AddSentryHook(dsn string) {
	sentryHook, err := NewSentryHook(SentryOptions{Dsn: dsn},
		logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel)
	if err != nil {
		Get().WithError(err).Error("unable to create new hook")
		return
	}
	defaultLog.Logger.AddHook(sentryHook)
}

// // AddFilenameHook adds hook that includes
// // filename and line number into the log.
// func AddFilenameHook() {
// 	filenameHook := filename.NewHook()
// 	filenameHook.Field = "file"
// 	Get().Logger.AddHook(filenameHook)
// }

// DefaultForRequest returns default logger with included http.Request details.
func DefaultForRequest(r *http.Request) *logrus.Entry {
	return IncludeRequest(defaultLog, r)
}

// IncludeRequest includes http.Request details into the log.Entry.
func IncludeRequest(log *logrus.Entry, r *http.Request) *logrus.Entry {
	reqID := middleware.GetReqID(r.Context())

	return log.
		WithFields(logrus.Fields{
			"req_id": reqID,
			"path":   r.URL.Path,
			"method": r.Method,
			"sender": r.Header.Get("X-Forwarded-For"),
		})
}
