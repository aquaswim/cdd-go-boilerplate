package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info"`
	LogPretty bool   `envconfig:"LOG_PRETTY" default:"false"`
}

func Get() *Config {
	var cfg Config
	envconfig.MustProcess("", &cfg)
	return &cfg
}
