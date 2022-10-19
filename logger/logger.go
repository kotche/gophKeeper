package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Init init logger
func Init(level string) *zerolog.Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &logger
}
