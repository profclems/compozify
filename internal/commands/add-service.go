package commands

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"github.com/profclems/compozify/pkg/parser"
)

type addServiceOpts struct {
	Logger *zerolog.Logger

	File        string
	Command     string
	Write       bool
	ServiceName string
}

func newAddServiceCmd(logger *zerolog.Logger) *cobra.Command {
	opts := addServiceOpts{
		Logger: logger,
	}
	cmd := &cobra.Command{
		Use:   "add-service [flags] DOCKER_RUN_COMMAND",
		Short: "Add a service to an existing docker-compose file",
		Long: `Converts the docker run command to docker compose and adds as a new service to an existing docker-compose file.
If no file is specified, compozify will look for a docker compose file in the current directory.
If no file is found, compozify will create one in the current directory.
Expected file names are docker-compose.[yml,yaml], compose.[yml,yaml]
`,
		Example: `
# add service to existing docker-compose file in current directory
$ compozify add-service "docker run -i -t --rm alpine"

# add service to existing docker-compose file
$ compozify add-service -f /path/to/docker-compose.yml "docker run -i -t --rm alpine"

# write to file
$ compozify add-service -w -f /path/to/docker-compose.yml "docker run -i -t --rm alpine"

# alternative usage specifying beginning of docker run command without quotes
$ compozify add-service -w -f /path/to/docker-compose.yml -- docker run -i -t --rm alpine

# add service with custom name
$ compozify add-service -w -f /path/to/docker-compose.yml -n my-service "docker run -i -t --rm alpine"
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}
			opts.Command = args[0]

			return addServiceRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.ServiceName, "service-name", "n", "", "Name of the service")
	cmd.Flags().BoolVarP(&opts.Write, "write", "w", false, "write to file")
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "Compose file path")

	return cmd
}

func addServiceRun(opts *addServiceOpts) error {
	readFile := opts.File != ""
	if opts.File == "" {
		opts.Logger.Info().Msg("No compose file specified. Searching for compose file in current directory")
		expectedFiles := []string{"docker-compose.yml", "docker-compose.yaml", "compose.yml", "compose.yaml"}

		for _, file := range expectedFiles {
			if _, err := os.Stat(file); err == nil {
				opts.File = file
				opts.Logger.Info().Msgf("Found compose file: %s", file)
				break
			}
		}

		if opts.File == "" {
			opts.Logger.Warn().Msg("No compose file found. Specify with --file or -f flag")
			opts.File = defaultFilename
			readFile = false
		}
	}

	var b []byte
	var err error

	if readFile {
		b, err = os.ReadFile(opts.File)
		if err != nil {
			return err
		}
	}

	p, err := parser.AppendToYAML(b, opts.Command)
	if err != nil {
		return err
	}

	if opts.ServiceName != "" {
		p.SetServiceName(opts.ServiceName)
	}

	err = p.Parse()
	if err != nil {
		return err
	}

	return printOutput(p, opts.Logger, opts.Write, opts.File)
}
