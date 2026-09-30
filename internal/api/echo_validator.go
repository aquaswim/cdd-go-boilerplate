package api

import (
	appErrors "cdd-go-boilerplate/internal/app_errors"

	"github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v5"
)

type echoValidator struct {
	validator *validator.Validate
}

func (e echoServer) newValidator() echo.Validator {
	return &echoValidator{
		validator: e.validate,
	}
}

func (e echoValidator) Validate(data any) error {
	err := e.validator.Struct(data)
	if err != nil {
		// note: use validator translator package
		//  payload format for error
		return appErrors.ErrTypeValidation.Wrap(err, "validation failed").
			WithProperty(errorx.PropertyPayload(), map[string]any{
				"details": err.Error(),
			})
	}
	return nil

}
