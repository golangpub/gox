package gox

import (
	"github.com/gopub/gox/internal"
	"net/http"
	"strings"
)

type Error interface {
	error
	Code() int
	Message() string
}

func NewError(code int, msg string) Error {
	if len(msg) == 0 {
		msg = http.StatusText(code)
	}
	return internal.NewError(code, msg)
}

func InternalError(message string) Error {
	return NewError(http.StatusInternalServerError, message)
}

func BadRequest(message string) Error {
	return NewError(http.StatusBadRequest, message)
}

func Unauthorized(message string) Error {
	return NewError(http.StatusUnauthorized, message)
}

func Forbidden(message string) Error {
	return NewError(http.StatusForbidden, message)
}

func NotFound(message string) Error {
	return NewError(http.StatusNotFound, message)
}

func Conflict(message string) Error {
	return NewError(http.StatusConflict, message)
}

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

type SubError interface {
	Error
	SubCode() int
}

func NewSubError(code, subCode int, msg string) Error {
	return internal.NewSubError(code, subCode, msg)
}

func InternalSubError(subCode int, message string) Error {
	return NewSubError(http.StatusInternalServerError, subCode, message)
}

func BadRequestSub(subCode int, message string) Error {
	return NewSubError(http.StatusBadRequest, subCode, message)
}

func UnauthorizedSub(subCode int, message string) Error {
	return NewSubError(http.StatusUnauthorized, subCode, message)
}

func ForbiddenSub(subCode int, message string) Error {
	return NewSubError(http.StatusForbidden, subCode, message)
}

func NotFoundFieldSub(subCode int, message string) Error {
	return NewSubError(http.StatusNotFound, subCode, message)
}

func ConflictSub(subCode int, message string) Error {
	return NewSubError(http.StatusConflict, subCode, message)
}
