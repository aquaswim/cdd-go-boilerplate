package api

import (
	"cdd-go-boilerplate/internal/entity"
	"cdd-go-boilerplate/internal/pkg/errorx"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"net/http"
)

func ErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		code := http.StatusInternalServerError
		response := entity.Error{
			Code:    "99",
			Error:   nil,
			Message: "Internal Server Error",
		}
		l := zerolog.Ctx(c.Request().Context())

		{
			var e *errorx.Errorx
			var echoError *echo.HTTPError
			if errors.As(err, &e) {
				code = translateErrCodeToHttpCode(e.Code)
				response.Error = e.Data
				response.Message = e.Message
				response.Code = e.Code
				response.Edited = e.Edited
			}
			if errors.As(err, &echoError) {
				l.Error().Err(err).Msg("Framework error received")
				response.Message = echoError.Error()
				response.Code = "FRAMEWORK"
			} else {
				// unknown error type will be printed to log
				// Echo's HttpError (like 404) is not handled yet and treated as Internal Error
				l.Error().Err(err).Type("errorType", err).Msg("Receive unknown error type while processing request")
			}
		}

		err2 := c.JSON(code, response)
		if err2 != nil {
			l.Error().AnErr("error", err).AnErr("sendError", err2).Msg("Error while sending error response")
		}
	}
}

func translateErrCodeToHttpCode(code string) int {
	switch code {
	case errorx.ErrCodeBadRequest:
		return http.StatusBadRequest
	case errorx.ErrCodeNotFound:
		return http.StatusNotFound
	case errorx.ErrCodeUnauthorized:
		return http.StatusUnauthorized
	case errorx.ErrCodeForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
