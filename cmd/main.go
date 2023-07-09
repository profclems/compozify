package main

import (
	"context"
	"github.com/profclems/composeify/internal/commands"
	"github.com/profclems/composeify/internal/version"
	"github.com/rs/zerolog"
)

func main() {
	// create logger
	logger := newLogger()

	cmd := commands.NewRootCmd(logger, version.GetVersion())
	err := cmd.ExecuteContext(context.Background())
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to run")
	}
}

func newLogger() *zerolog.Logger {
	out := zerolog.NewConsoleWriter()
	logger := zerolog.New(out)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &logger
}
