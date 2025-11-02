package validation

import "github.com/go-playground/validator/v10"

type ErrorMetadata struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value any    `json:"value"`
}

func FormatError(err error) []ErrorMetadata {
	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}
	res := make([]ErrorMetadata, len(validationErrs))
	for i, err := range validationErrs {
		res[i] = ErrorMetadata{
			Field: err.Field(),
			Tag:   err.Tag(),
			Value: err.Value(),
		}
	}
	return res
}
