package main

import (
	"cdd-go-boilerplate/internal"
	"cdd-go-boilerplate/internal/api"
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func main() {
	c := internal.InitContainer()

	server := internal.Resolve[api.Server](c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer func() {
		stop()
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
