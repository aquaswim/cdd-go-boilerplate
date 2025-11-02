package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"
	"cdd-go-boilerplate/internal/entity"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func validateStruct(v *validator.Validate, data any) error {
	err := v.Struct(data)
	if err != nil {
		// note: use validator translator package
		return appErrors.ErrTypeValidation.Wrap(err, "validation failed")
	}
	return nil
}

func bindAndValidate[T any](ctx echo.Context, v *validator.Validate) (*T, error) {
	data := new(T)
	if err := ctx.Bind(data); err != nil {
		return nil, appErrors.ErrTypeBind.Wrap(err, "failed to bind request")
	}
	if err := validateStruct(v, data); err != nil {
		return nil, err
	}
	return data, nil
}

func sendSuccessResponse(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, entity.NewSuccessResponse(data))
}
