package validation

import "github.com/go-playground/validator/v10"

func NewValidator() *validator.Validate {
	// note: register custom validation here if needed
	return validator.New(validator.WithRequiredStructEnabled())
}
