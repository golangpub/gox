package types

import (
	"fmt"
)

const InternalErrorCode = 99999

type Error interface {
	error
	Code() int
	Message() string
}

// errorInfo represents error info
type errorInfo struct {
	code    int
	message string
}

func (e *errorInfo) Code() int {
	if e == nil {
		return 0
	}

	return e.code
}

func (e *errorInfo) Message() string {
	if e == nil {
		return ""
	}

	return e.message
}

func (e *errorInfo) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code:%d message:%s", e.code, e.message)
}

func NewError(code int, msg string) Error {
	if code <= 0 || code > InternalErrorCode {
		panic(fmt.Sprintf("code value should be (0, %d]", code))
	}
	return &errorInfo{code: code, message: msg}
}

func InternalError(message string) Error {
	return NewError(InternalErrorCode, message)
}
