package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/echotest"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorHandler(t *testing.T) {
	log.Logger = log.Output(zerolog.NewConsoleWriter(zerolog.ConsoleTestWriter(t)))

	fn := ErrorHandler()

	req, err := http.NewRequestWithContext(log.Logger.WithContext(t.Context()), http.MethodGet, "/", nil)
	require.NoError(t, err)

	t.Run("already commited response", func(t *testing.T) {
		ctx, rec := echotest.ContextConfig{
			Request: req,
		}.ToContextRecorder(t)

		require.NoError(t, ctx.JSON(http.StatusTeapot, nil))

		fn(ctx, appErrors.ErrTypeInternal.NewWithNoMessage())

		assert.Equal(t, http.StatusTeapot, rec.Code)
	})

	t.Run("test app error", func(t *testing.T) {
		ctx, rec := echotest.ContextConfig{
			Request: req,
		}.ToContextRecorder(t)

		fn(ctx, appErrors.ErrTypeNotFound.New("test error"))

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("test framework error", func(t *testing.T) {
		ctx, rec := echotest.ContextConfig{
			Request: req,
		}.ToContextRecorder(t)

		fn(ctx, echo.ErrNotFound)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("test unknown error", func(t *testing.T) {
		ctx, rec := echotest.ContextConfig{
			Request: req,
		}.ToContextRecorder(t)

		fn(ctx, errors.New("test error"))

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
