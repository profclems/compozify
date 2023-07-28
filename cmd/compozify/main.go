package main

import (
	"context"
	"os"

	"github.com/profclems/compozify/internal/commands"
	"github.com/profclems/compozify/internal/utils"
	"github.com/profclems/compozify/internal/version"
)

func main() {
	// create logger
	logger := utils.NewLogger()

	cmd := commands.NewRootCmd(logger, version.GetVersion())
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	if err := cmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
