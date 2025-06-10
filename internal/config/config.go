package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info"`
	LogPretty bool   `envconfig:"LOG_PRETTY" default:"false"`

	DBPgHost     string `envconfig:"DB_PG_ADDR" required:"true"`
	DBPgUsername string `envconfig:"DB_PG_USER" required:"true"`
	DBPgDBName   string `envconfig:"DB_PG_DBNAME" required:"true"`
	DBPgPassword string `envconfig:"DB_PG_PASS" required:"true"`
}

func Get() *Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return &cfg
}
