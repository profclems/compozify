package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/profclems/compozify/pkg/parser"
)

var defaultFilename = "compose.yml"

type convertOpts struct {
	Command     string
	Write       bool
	OutFilePath string

	Logger *zerolog.Logger
}

func newConvertCmd(logger *zerolog.Logger) *cobra.Command {
	opts := convertOpts{
		Logger: logger,
	}

	cmd := &cobra.Command{
		Use:   "convert [flags] DOCKER_RUN_COMMAND",
		Short: "convert docker run command to docker compose file",
		Example: `
# convert and write to stdout
$ compozify convert "docker run -i -t --rm alpine"

# write to file
$ compozify convert -w "docker run -i -t --rm alpine"

# write to file with custom name
$ compozify convert -w -o docker-compose.yml "docker run -i -t --rm alpine"

# alternative usage specifying beginning of docker run command
$ compozify convert -w -- docker run -i -t --rm alpine
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			switch len(args) {
			case 0:
				return cmd.Help()
			case 1:
				opts.Command = args[0]
			default:
				opts.Command = strings.Join(args, " ")
			}

			return convertRun(&opts)
		},
		Args: cobra.MinimumNArgs(1),
	}

	cmd.Flags().BoolVarP(&opts.Write, "write", "w", false, "write to file")
	cmd.Flags().StringVarP(&opts.OutFilePath, "out", "o", defaultFilename, "output file path")

	return cmd
}

func convertRun(opts *convertOpts) (err error) {
	log := opts.Logger

	log.Info().Msg("Parsing Docker run command")
	log.Debug().Msgf("Docker run command: %s", opts.Command)

	parser, err := parser.New(opts.Command)
	if err != nil {
		return err
	}

	log.Info().Msg("Generating Docker compose file")
	err = parser.Parse()
	if err != nil {
		return err
	}
	log.Info().Msg("Docker compose file generated")

	writer := os.Stdout
	if opts.Write {
		log.Info().Msgf("Writing to file %s", opts.OutFilePath)
		writer, err = os.Create(opts.OutFilePath)
		if err != nil {
			return err
		}
		defer func(writer *os.File) {
			e := writer.Close()
			if e != nil {
				log.Error().Err(e).Msg("Error closing file")
			}
			err = e
		}(writer)
	}
	fmt.Fprintf(writer, "%s", parser.String())

	return err
}
