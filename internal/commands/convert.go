package commands

import (
	"fmt"
	"github.com/profclems/composeify/pkg/parser"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
	"strings"
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

func convertRun(opts *convertOpts) error {
	parser, err := parser.NewParser(opts.Command)
	if err != nil {
		return err
	}

	err = parser.Parse()
	if err != nil {
		return err
	}

	writer := os.Stdout
	if opts.Write {
		writer, err = os.Create(opts.OutFilePath)
		if err != nil {
			return err
		}
		defer writer.Close()
	}

	fmt.Fprintf(writer, "%s", parser.String())

	return nil
}
