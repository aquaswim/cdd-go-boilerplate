package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"
	"cdd-go-boilerplate/internal/entity"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/rs/zerolog"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(c *echo.Context, err error) {
		l := zerolog.Ctx(c.Request().Context())
		resp, err2 := echo.UnwrapResponse(c.Response())
		if err2 != nil {
			l.Error().Err(err2).Msg("failed to unwrap response")
			return
		}

		if resp.Committed {
			return
		}
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
		} else if statusCode := echo.StatusCode(err); statusCode != 0 {
			// handle echo error
			l.Error().Err(err).Msg("framework error detected")
			httpCode = statusCode
			errResp.Code = "FRAMEWORK"
			errResp.Message = fmt.Sprint(err)
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
