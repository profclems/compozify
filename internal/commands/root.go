package commands

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/profclems/composeify/internal/version"
)

func NewRootCmd(logger *zerolog.Logger, version version.Info) *cobra.Command {
	var verbose bool
	cmd := &cobra.Command{
		Use:   "compozify",
		Short: "compozify is a tool mainly for converting docker run commands to docker compose files",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if verbose {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}
			return nil
		},
	}

	cmd.Version = version.String()

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", version.IsDev(), "verbose output")

	cmd.AddCommand(newConvertCmd(logger))

	return cmd
}
