package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"

	"github.com/profclems/compozify/internal/commands"
	"github.com/profclems/compozify/internal/version"
)

func main() {
	var flagErr pflag.ErrorHandling
	docsCmd := pflag.NewFlagSet("", flagErr)
	manpage := docsCmd.BoolP("manpage", "m", false, "Generate manual pages instead of web docs")
	path := docsCmd.StringP("path", "p", "./docs/source/", "Path where you want the generated docs saved")
	help := docsCmd.BoolP("help", "h", false, "Help about any command")

	if err := docsCmd.Parse(os.Args); err != nil {
		fatal(err)
	}
	if *help {
		_, err := fmt.Fprintf(os.Stderr, "Usage of %s:\n\n%s", os.Args[0], docsCmd.FlagUsages())
		if err != nil {
			fatal(err)
		}
		os.Exit(1)
	}
	err := os.MkdirAll(*path, 0o755)
	if err != nil {
		fatal(err)
	}

	logger := zerolog.Nop()
	cmd := commands.NewRootCmd(&logger, version.GetVersion())
	cmd.DisableAutoGenTag = true
	if *manpage {
		if err := genManPage(cmd, *path); err != nil {
			fatal(err)
		}
	} else {
		err := doc.GenMarkdownTree(cmd, *path)
		if err != nil {
			fatal(err)
		}
	}
}

func genManPage(cmd *cobra.Command, path string) error {
	header := &doc.GenManHeader{
		Title:   "compozify",
		Section: "1",
		Source:  "",
		Manual:  "",
	}
	err := doc.GenManTree(cmd, header, path)
	if err != nil {
		return err
	}
	return nil
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
