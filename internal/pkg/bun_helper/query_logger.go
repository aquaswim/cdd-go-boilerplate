package bunHelper

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

const slowQueryThreshold = 3000 * time.Millisecond

type queryLogger struct {
}

func newQueryLogger() bun.QueryHook {
	return &queryLogger{}
}

func (q queryLogger) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context {
	return ctx
}

func (q queryLogger) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	zerolog.Ctx(ctx).WithLevel(getLogLevel(event)).
		Err(event.Err).
		Dur("duration", time.Since(event.StartTime)).
		Msgf("DB Query: %s", event.Query)
}

func getLogLevel(event *bun.QueryEvent) zerolog.Level {
	if event.Err != nil {
		return zerolog.ErrorLevel
	}
	if time.Since(event.StartTime) > slowQueryThreshold {
		return zerolog.WarnLevel
	}
	return zerolog.InfoLevel
}
