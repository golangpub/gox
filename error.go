package gox

import (
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorString string

func (es ErrorString) Error() string {
	return string(es)
}

const (
	ErrNoValue ErrorString = "no value"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	if len(message) == 0 {
		message = http.StatusText(code % 1000)
	}
	return &Error{
		Code:    code,
		Message: message,
	}
}

func InternalError(message string) *Error {
	return NewError(http.StatusInternalServerError, message)
}

func BadRequest(message string) *Error {
	return NewError(http.StatusBadRequest, message)
}

func Unauthorized(message string) *Error {
	return NewError(http.StatusUnauthorized, message)
}

func Forbidden(message string) *Error {
	return NewError(http.StatusForbidden, message)
}

func NotFound(message string) *Error {
	return NewError(http.StatusNotFound, message)
}

func Conflict(message string) *Error {
	return NewError(http.StatusConflict, message)
}

func ToStatusError(err error) error {
	if err == nil {
		return nil
	}

	// if err is status error, return directly
	_, ok := status.FromError(err)
	if ok {
		return err
	}

	err = errors.Cause(err)

	if err == ErrNoValue {
		return status.Error(codes.Code(http.StatusNotFound), err.Error())
	}

	switch v := err.(type) {
	case *Error:
		return status.Error(codes.Code(v.Code), v.Error())
	default:
		return status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}
}

func FromStatusError(err error) *Error {
	if err == nil {
		return nil
	}

	s, ok := status.FromError(err)
	if !ok {
		return InternalError(err.Error())
	}

	return NewError(int(s.Code()), s.Message())
}
