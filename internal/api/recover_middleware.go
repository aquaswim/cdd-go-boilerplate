package api

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

// RecoverMiddleware is echo recover middleware with global logger integration
func RecoverMiddleware() echo.MiddlewareFunc {
	return echoMiddleware.RecoverWithConfig(echoMiddleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			zerolog.Ctx(c.Request().Context()).Error().Err(err).Str("stack", string(stack)).Msg("Panic Recovered")
			return err
		},
	})
}
