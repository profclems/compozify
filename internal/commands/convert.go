package commands

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/profclems/compozify/pkg/parser"
)

var defaultFilename = "compose.yml"

type convertOpts struct {
	Command       string
	OutFilePath   string
	Write         bool
	AppendService bool

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

			if opts.AppendService && opts.OutFilePath == "" {
				return fmt.Errorf("--append-service requires --out flag")
			}

			return convertRun(&opts)
		},
		Args: cobra.MinimumNArgs(1),
	}

	cmd.Flags().BoolVarP(&opts.AppendService, "append-service", "a", false, "append service to existing compose file. Requires --out flag")
	cmd.Flags().BoolVarP(&opts.Write, "write", "w", false, "write to file")
	cmd.Flags().StringVarP(&opts.OutFilePath, "out", "o", defaultFilename, "output file path")

	return cmd
}

func convertRun(opts *convertOpts) (err error) {
	log := opts.Logger

	log.Info().Msg("Parsing Docker run command")
	log.Debug().Msgf("Docker run command: %s", opts.Command)

	if opts.AppendService {
		log.Info().Msg("Appending service to existing compose file")
		return addServiceRun(&addServiceOpts{
			Logger:  log,
			File:    opts.OutFilePath,
			Command: opts.Command,
			Write:   opts.Write,
		})
	}

	p, err := parser.New(opts.Command)
	if err != nil {
		return err
	}

	log.Info().Msg("Generating Docker compose file")
	err = p.Parse()
	if err != nil {
		return err
	}
	log.Info().Msg("Docker compose file generated")

	return printOutput(p, log, opts.Write, opts.OutFilePath)
}
