package internal

import (
	"encoding/json"
	"fmt"
)

type fieldErrorJSONObject struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Field   string `json:"field"`
}

// FieldError represents error info
type FieldError struct {
	code    int
	message string
	field   string
}

func NewFieldError(code int, msg, field string) *FieldError {
	return &FieldError{code: code, message: msg, field: field}
}

func (e *FieldError) UnmarshalJSON(b []byte) error {
	var obj *fieldErrorJSONObject
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	e.code = obj.Code
	e.message = obj.Message
	e.field = obj.Field
	return nil
}

func (e *FieldError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&fieldErrorJSONObject{
		Code:    e.code,
		Message: e.message,
		Field:   e.field,
	})
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
	return fmt.Sprintf("%d:%s {%s}", e.code, e.message, e.field)
}
