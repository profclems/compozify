package main

import (
	"context"
	"github.com/profclems/composeify/internal/commands"
	"github.com/profclems/composeify/internal/version"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func main() {
	// create logger
	logger := newLogger()

	cmd := commands.NewRootCmd(logger, version.GetVersion())
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	if err := cmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}

func newLogger() *zerolog.Logger {
	out := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.TimeFormat = time.RFC3339
	})

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Output(out)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &logger
}
