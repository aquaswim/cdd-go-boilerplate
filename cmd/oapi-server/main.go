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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := server.Start(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}

	log.Info().Msg("shutting down server")
}
