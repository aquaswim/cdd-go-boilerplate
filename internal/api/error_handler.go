package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		l := zerolog.Ctx(c.Request().Context())
		errx := errorx.Cast(err)
		if echoErr, ok := err.(*echo.HTTPError); ok {
			// framework error
			errx = errorx.Decorate(
				appErrors.FrameworkError.WithProperty(appErrors.HttpCodeProperty, echoErr.Code),
				"%s",
				echoErr.Message,
			).WithUnderlyingErrors(err)
		} else if errx == nil {
			// unknown error
			l.Error().Err(err).Msg("unknown error occurred")
			errx = errorx.Decorate(appErrors.FrameworkError, "%s", err).WithUnderlyingErrors(err)
		}

		l.Error().Err(errx).Stack().Msg("error occurred")

		errResp, httpCode := appErrors.ExtractAppError(errx)

		err = c.JSON(httpCode, errResp)
		if err != nil {
			l.Error().
				Err(err).
				Int("code", httpCode).
				Any("response", errResp).
				Msg("failed to send error response")
		}
	}
}
