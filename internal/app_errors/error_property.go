package appErrors

import (
	"cdd-go-boilerplate/internal/entity"

	"github.com/joomcode/errorx"
)

var EditedProperty = errorx.RegisterProperty("edited")

func extractPropertyTo[T any](err error, prop errorx.Property, target *T) {
	val, ok := errorx.ExtractProperty(err, prop)
	if !ok {
		return
	}

	if v, ok := val.(T); ok {
		*target = v
	}
}

func ExtractAppError(rawErr error) (*entity.Error, int, bool) {
	err := errorx.Cast(rawErr)
	if err == nil {
		return nil, 0, false
	}

	httpCode := typeToHttpCode(err)
	out := new(entity.Error)
	out.Message = err.Message()

	extractPropertyTo(err, EditedProperty, &out.Edited)
	extractPropertyTo(err, errorx.PropertyPayload(), &out.Error)
	out.Code = err.Type().FullName()

	return out, httpCode, true
}
