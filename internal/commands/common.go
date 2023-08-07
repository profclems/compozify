package commands

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/profclems/compozify/pkg/parser"
)

func printOutput(parser *parser.Parser, log *zerolog.Logger, writeToFile bool, path string) error {
	writer := os.Stdout
	var err error
	if writeToFile {
		log.Info().Msgf("Writing to file %s", path)
		writer, err = os.Create(path)
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
	_, err = fmt.Fprintf(writer, "%s", parser.String())
	return err
}
