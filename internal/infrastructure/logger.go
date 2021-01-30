package infrastructure

import (
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	env    string
	Logger *logrus.Logger
}

func NewLogger(env, logLevel, logFormat string) *LogrusLogger {
	logger := logrus.New()
	l, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithField("level", logLevel).Warn("Invalid log level, fallback to 'info'")
	} else {
		logrus.SetLevel(l)
	}

	switch logFormat {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	return &LogrusLogger{env: env, Logger: logger}
}

func (l LogrusLogger) Log(args ...interface{}) {
	if l.Logger == nil {
		return
	}

	if len(args) == 2 {
		castedError, ok := args[0].(error)
		if ok {
			l.newLog(castedError, args[1])
		} else {
			l.Logger.Info(args...)
		}
		return
	}

	l.Logger.Info(args...)
}

func (l LogrusLogger) newLog(err error, usecase interface{}) {
	l.Logger.WithError(err).Error(usecase)
}
