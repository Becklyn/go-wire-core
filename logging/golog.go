package logging

import (
	"io"
	"os"
	"time"

	"github.com/Becklyn/go-wire-core/env"
	"github.com/fraym/golog"
	"github.com/rs/zerolog"
)

var (
	LOG_LEVEL  = "LOG_LEVEL"
	LOG_FORMAT = "LOG_FORMAT"
)

func New(_ *env.Env) golog.Logger {
	logger := golog.NewZerologLogger()

	logger.SetLevel(getLogLevel())
	logger.SetOutput(getLogFormatOutput())
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

func getLogFormatOutput() io.Writer {
	switch env.String(LOG_FORMAT) {
	case "json":
		return os.Stdout
	default:
		return &zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}
	}
}
