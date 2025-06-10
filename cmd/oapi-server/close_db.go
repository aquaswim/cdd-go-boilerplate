package main

import (
	"github.com/golobby/container/v3"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
)

func closesDB(c container.Container) {
	log.Info().Msg("Closing database connection")
	container.MustCall(c, func(db *bun.DB) {
		_closeDB(db, "main")
	})
}

func _closeDB(db *bun.DB, name string) {
	err := db.Close()
	if err != nil {
		log.Error().Err(err).Str("name", name).Msg("failed to close db")
		return
	}
	log.Info().Str("name", name).Msg("db closed")
}
