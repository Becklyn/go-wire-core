package logging

import (
	"github.com/Becklyn/go-wire-core/env"
	"github.com/fraym/golog"
)

var LOG_LEVEL = "LOG_LEVEL"

func New(_ *env.Env) golog.Logger {
	logger := golog.NewZerologLogger()

	logger.SetLevel(getLogLevel())
	logger.Info().WithField("environment", env.String(env.APP_ENV)).Write("Using environment")

	return logger
}

func getLogLevel() golog.Level {
	switch env.String(LOG_LEVEL) {
	case "debug":
		return golog.DebugLevel
	case "warn":
		return golog.WarnLevel
	case "error":
		return golog.ErrorLevel
	case "fatal":
		return golog.FatalLevel
	default:
		return golog.InfoLevel
	}
}
