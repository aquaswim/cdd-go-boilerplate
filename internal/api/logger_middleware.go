package api

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func LoggerMiddleware(logger *zerolog.Logger) echo.MiddlewareFunc {
	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		//LogStatus:       true, // this is not working since the status set in error handler
		LogLatency:      true,
		LogRemoteIP:     true,
		LogMethod:       true,
		LogURI:          true,
		LogRequestID:    true,
		LogProtocol:     true,
		LogUserAgent:    true,
		LogResponseSize: true,
		LogError:        true,
		BeforeNextFunc: func(c echo.Context) {
			// inject logger to context
			lCtx := logger.
				With().
				Str("reqId", c.Response().Header().Get(echo.HeaderXRequestID)). // reqid from echo's RequestID middleware
				Logger().
				WithContext(c.Request().Context())
			c.SetRequest(c.Request().WithContext(lCtx))
		},
		LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
			// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#HttpRequest
			logDetail := zerolog.Dict().
				Str("requestMethod", v.Method).
				Str("requestUrl", v.URI).
				Str("protocol", v.Protocol).
				Str("userAgent", v.UserAgent).
				Int64("responseSize", v.ResponseSize).
				Dur("latency", v.Latency)
			var level = zerolog.InfoLevel
			if v.Error != nil {
				level = zerolog.ErrorLevel
			}
			// msg format: HTTPREQ: LATENCY | IP | METHOD PATH - REQ ID
			zerolog.Ctx(c.Request().Context()).
				WithLevel(level).
				Err(v.Error).
				Dict("httpRequest", logDetail).
				Msgf("HTTPREQ: %s | %s | %s %s - %s", v.Latency, v.RemoteIP, v.Method, v.URI, v.RequestID)
			return nil
		},
	})
}
