package appErrors

import (
	"github.com/joomcode/errorx"
)

var appErrorNs = errorx.NewNamespace("app")

// define all error types here
var (
	ErrTypeInternal     = appErrorNs.NewType("internal")
	ErrTypeValidation   = appErrorNs.NewType("validation")
	ErrTypeBind         = ErrTypeValidation.NewSubtype("bind")
	ErrTypeNotFound     = appErrorNs.NewType("not_found")
	ErrTypeUnauthorized = appErrorNs.NewType("unauthorized")
	ErrTypeForbidden    = appErrorNs.NewType("forbidden")
)

func typeToHttpCode(err error) int {
	switch errorx.TypeSwitch(err,
		ErrTypeNotFound,
		ErrTypeValidation,
		ErrTypeInternal,
	) {
	case ErrTypeNotFound:
		return 404
	case ErrTypeValidation:
		return 400
	default:
		return 500
	}
}
