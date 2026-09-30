package globalLogger

import (
	"log/slog"

	"github.com/rs/zerolog/log"
	slogzerolog "github.com/samber/slog-zerolog/v2"
)

func SlogAdapter() *slog.Logger {
	return slog.New(slogzerolog.Option{
		Logger: &log.Logger,
	}.NewZerologHandler())
}
