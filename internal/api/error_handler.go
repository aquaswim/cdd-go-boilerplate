package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"
	"cdd-go-boilerplate/internal/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		l := zerolog.Ctx(c.Request().Context())
		l.Error().Msgf("Error detected: %+v", err)

		httpCode := http.StatusInternalServerError
		errResp := entity.Error{
			Code:    appErrors.ErrTypeInternal.String(),
			Edited:  false,
			Error:   nil,
			Message: "Internal server error.",
		}

		if resp, code, ok := appErrors.ExtractAppError(err); ok {
			l.Error().Err(err).Msgf("Error detected: %+v", resp)
			httpCode = code
			errResp = *resp
		} else if echoErr, ok := err.(*echo.HTTPError); ok {
			// handle echo error
			l.Error().Err(echoErr).Msg("framework error detected")
			httpCode = echoErr.Code
			errResp.Code = "FRAMEWORK"
			errResp.Message = fmt.Sprint(echoErr.Message)
		} else {
			l.Error().Err(err).Msg("unknown error detected")
		}

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
