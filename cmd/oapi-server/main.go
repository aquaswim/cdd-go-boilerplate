package main

import (
	"cdd-go-boilerplate/internal"
	"cdd-go-boilerplate/internal/api"
	"cdd-go-boilerplate/internal/pkg/utils"
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

func main() {
	c := internal.InitContainer()

	server := utils.Resolve[api.Server](c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer func() {
		stop()
		closesDB(c)
	}()

	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()
	<-ctx.Done()
	log.Info().Msg("shutting down server")
	err := server.Stop()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to stop server")
	}
}
