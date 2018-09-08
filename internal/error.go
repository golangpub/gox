package internal

import "fmt"

// Error represents error info
type Error struct {
	code    int
	message string
}

func NewError(code int, msg string) *Error {
	return &Error{code: code, message: msg}
}

func (e *Error) Code() int {
	if e == nil {
		return 0
	}

	return e.code
}

func (e *Error) Message() string {
	if e == nil {
		return ""
	}

	return e.message
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code:%d message:%s", e.code, e.message)
}
