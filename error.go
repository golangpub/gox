package types

import (
	"github.com/gopub/types/internal"
)

const InternalErrorCode = 99999

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
	return NewError(InternalErrorCode, message)
}
