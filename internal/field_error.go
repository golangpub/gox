package internal

import "fmt"

// FieldError represents error info
type FieldError struct {
	code    int
	message string
	field   string
}

func NewFieldError(code int, msg, field string) *FieldError {
	return &FieldError{code: code, message: msg, field: field}
}

func (e *FieldError) Code() int {
	if e == nil {
		return 0
	}

	return e.code
}

func (e *FieldError) Message() string {
	if e == nil {
		return ""
	}

	return e.message
}

func (e *FieldError) Field() string {
	if e == nil {
		return ""
	}

	return e.field
}

func (e *FieldError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code:%d message:%s field:%s", e.code, e.message, e.field)
}
