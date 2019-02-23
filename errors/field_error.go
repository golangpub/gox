package errors

import (
	"github.com/gopub/gox/errors/internal"
	"net/http"
	"strings"
)

type FieldError interface {
	Error
	Field() string
}

func NewFieldError(code int, msg, field string) Error {
	return internal.NewFieldError(code, msg, field)
}

func BadRequestField(field string) Error {
	return NewFieldError(http.StatusBadRequest, strings.ToLower(http.StatusText(http.StatusBadRequest)), field)
}

func UnauthorizedField(field string) Error {
	return NewFieldError(http.StatusUnauthorized, strings.ToLower(http.StatusText(http.StatusUnauthorized)), field)
}

func ForbiddenField(field string) Error {
	return NewFieldError(http.StatusForbidden, strings.ToLower(http.StatusText(http.StatusForbidden)), field)
}

func NotFoundField(field string) Error {
	return NewFieldError(http.StatusNotFound, strings.ToLower(http.StatusText(http.StatusNotFound)), field)
}

func ConflictField(field string) Error {
	return NewFieldError(http.StatusConflict, strings.ToLower(http.StatusText(http.StatusConflict)), field)
}
