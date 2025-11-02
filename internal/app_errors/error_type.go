package appErrors

import (
	"cdd-go-boilerplate/internal/entity"

	"github.com/joomcode/errorx"
)

var appErrorNs = errorx.NewNamespace("app")
var appErrorType = errorx.NewType(appErrorNs, "appError")

var HttpCodeProperty = errorx.RegisterProperty("httpCode")
var CodeProperty = errorx.RegisterProperty("code")
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

func ExtractAppError(err *errorx.Error) (*entity.Error, int) {
	httpCode := 500
	out := new(entity.Error)
	out.Message = err.Message()

	extractPropertyTo(err, HttpCodeProperty, &httpCode)
	extractPropertyTo(err, EditedProperty, &out.Edited)
	extractPropertyTo(err, CodeProperty, &out.Code)
	extractPropertyTo(err, errorx.PropertyPayload(), &out.Error)

	return out, httpCode
}
