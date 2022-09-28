package logging

import (
	"os"
	"time"

	"github.com/Becklyn/go-wire-core/env"
	"github.com/sirupsen/logrus"
)

var LOG_LEVEL = "LOG_LEVEL"

func New(_ *env.Env) *logrus.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC822,
	})
	logger.SetLevel(getLogLevel())
	logger.WithFields(logrus.Fields{
		"environment": env.String(env.APP_ENV),
	}).Info("Using environment")

	return logger
}

func getLogLevel() logrus.Level {
	switch env.String(LOG_LEVEL) {
	case "debug":
		return logrus.DebugLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
