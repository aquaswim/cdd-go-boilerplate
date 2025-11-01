package bunHelper

import (
	"database/sql"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type ConfigPg struct {
	Host     string
	DBName   string
	Username string
	Password string
}

func ConnectPg(cfg *ConfigPg) *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr(cfg.Host),
		pgdriver.WithDatabase(cfg.DBName),
		pgdriver.WithUser(cfg.Username),
		pgdriver.WithPassword(cfg.Password),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
		pgdriver.WithApplicationName("raya-faq"),
		// todo: remove this
		pgdriver.WithInsecure(true),
	))

	db := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())

	// adding query logger
	db.AddQueryHook(newQueryLogger())

	log.Info().Msgf("connecting to postgres %s", cfg.Host)
	err := db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to ping postgres")
	}

	log.Info().Msgf("connected to postgres %s", cfg.Host)
	return db
}
