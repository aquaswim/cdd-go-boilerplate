package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"
	"cdd-go-boilerplate/internal/entity"
	"net/http"

	"github.com/labstack/echo/v5"
)

func bindAndValidate[T any](ctx *echo.Context) (*T, error) {
	data := new(T)
	if err := ctx.Bind(data); err != nil {
		return nil, appErrors.ErrTypeBind.Wrap(err, "failed to bind request")
	}
	if err := ctx.Validate(data); err != nil {
		return nil, err
	}
	return data, nil
}

func sendSuccessResponse(ctx *echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, entity.NewSuccessResponse(data))
}
