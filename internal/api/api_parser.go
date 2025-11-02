package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"
	"cdd-go-boilerplate/internal/entity"
	"cdd-go-boilerplate/internal/pkg/validation"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func validateStruct(v *validator.Validate, data any) error {
	err := v.Struct(data)
	if err != nil {
		// note: use validator translator package
		return errorx.WithPayload(appErrors.ValidationError, validation.FormatError(err))
	}
	return nil
}

func bindAndValidate[T any](ctx echo.Context, v *validator.Validate) (*T, error) {
	data := new(T)
	if err := ctx.Bind(data); err != nil {
		return nil, appErrors.InternalError.WithUnderlyingErrors(err)
	}
	if err := validateStruct(v, data); err != nil {
		return nil, err
	}
	return data, nil
}

func sendSuccessResponse(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, entity.NewSuccessResponse(data))
}
