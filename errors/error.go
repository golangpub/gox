package errors

import (
	"github.com/gopub/types/errors/internal"
	"net/http"
)

type Error interface {
	error
	Code() int
	Message() string
}

func NewError(code int, msg string) Error {
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
