package appErrors

import (
	"net/http"

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
		ErrTypeInternal,
		ErrTypeValidation,
		ErrTypeBind,
		ErrTypeNotFound,
		ErrTypeUnauthorized,
		ErrTypeForbidden,
	) {
	case ErrTypeNotFound:
		return http.StatusNotFound
	case ErrTypeValidation, ErrTypeBind:
		return http.StatusBadRequest
	case ErrTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrTypeForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
