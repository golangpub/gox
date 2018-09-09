package errors

import (
	"github.com/gopub/types/errors/internal"
	"net/http"
	"strings"
)

type Error interface {
	error
	Code() int
	Message() string
}

type FieldError interface {
	Error
	Field() string
}

func NewError(code int, msg string) Error {
	return internal.NewError(code, msg)
}

func NewFieldError(code int, msg, field string) Error {
	return internal.NewFieldError(code, msg, field)
}

func InternalError(message string) Error {
	return NewError(http.StatusInternalServerError, message)
}

func BadRequest(message string) Error {
	return NewError(http.StatusBadRequest, message)
}

func BadRequestField(field string) Error {
	return NewFieldError(http.StatusBadRequest, strings.ToLower(http.StatusText(http.StatusBadRequest)), field)
}

func Unauthorized(message string) Error {
	return NewError(http.StatusUnauthorized, message)
}

func UnauthorizedField(field string) Error {
	return NewFieldError(http.StatusUnauthorized, strings.ToLower(http.StatusText(http.StatusUnauthorized)), field)
}

func Forbidden(message string) Error {
	return NewError(http.StatusForbidden, message)
}

func ForbiddenField(field string) Error {
	return NewFieldError(http.StatusForbidden, strings.ToLower(http.StatusText(http.StatusForbidden)), field)
}

func NotFound(message string) Error {
	return NewError(http.StatusNotFound, message)
}

func NotFoundField(field string) Error {
	return NewFieldError(http.StatusNotFound, strings.ToLower(http.StatusText(http.StatusNotFound)), field)
}

func Conflict(message string) Error {
	return NewError(http.StatusConflict, message)
}

func ConflictField(field string) Error {
	return NewFieldError(http.StatusConflict, strings.ToLower(http.StatusText(http.StatusConflict)), field)
}
