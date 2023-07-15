package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// NewLogger creates a new logger.
func NewLogger() *zerolog.Logger {
	out := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.TimeFormat = time.RFC3339
	})

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(out)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &logger
}
