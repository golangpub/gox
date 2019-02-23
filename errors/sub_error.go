package errors

import (
	"github.com/gopub/gox/errors/internal"
	"net/http"
)

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
