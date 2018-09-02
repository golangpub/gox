package types

import (
	"fmt"
)

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
	return &errorInfo{code: code, message: msg}
}
