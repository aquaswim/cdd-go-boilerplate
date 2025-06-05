package api

import (
	"cdd-go-boilerplate/internal/entity"
	"cdd-go-boilerplate/internal/pkg/errorx"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

func validateStruct(v *validator.Validate, data any) error {
	err := v.Struct(data)
	if err != nil {
		return errorx.New(errorx.ErrCodeBadRequest, err.Error())
	}
	return nil
}

func bindAndValidate[T any](ctx echo.Context, v *validator.Validate) (*T, error) {
	data := new(T)
	if err := ctx.Bind(data); err != nil {
		return nil, errorx.New(errorx.ErrCodeBadRequest, err.Error())
	}
	if err := validateStruct(v, data); err != nil {
		return nil, err
	}
	return data, nil
}

func sendSuccessResponse(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, entity.NewSuccessResponse(data))
}
