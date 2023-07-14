package main

import (
	"context"
	"flag"
	"github.com/profclems/compozify/api"
	"github.com/profclems/compozify/internal/utils"
	"github.com/profclems/compozify/web"
	"net"
)

var serverAddr = flag.String("addr", ":8080", "server address of the api gateway and frontend app")

func main() {
	// parse flags
	flag.Parse()
	// create logger
	logger := utils.NewLogger()

	listener, err := net.Listen("tcp", *serverAddr)
	if err != nil {
		logger.Fatal().Err(err).Str("address", *serverAddr).Msg("failed to listen to address")
	}

	logger.Info().Str("address", *serverAddr).Msg("listening on address")

	server := api.NewServer(logger, listener, web.Assets)

	ctx := context.Background()
	err = server.Run(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to run server")
	}

	defer func(server *api.Server) {
		err := server.Close()
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to close server")
		}
	}(server)
}
