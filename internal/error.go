package internal

import (
	"encoding/json"
)

type errorJSONObject struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error represents error info
type Error struct {
	code    int
	message string
}

func NewError(code int, msg string) *Error {
	return &Error{code: code, message: msg}
}

func (e *Error) UnmarshalJSON(b []byte) error {
	var obj *errorJSONObject
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	e.code = obj.Code
	e.message = obj.Message
	return nil
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&errorJSONObject{
		Code:    e.code,
		Message: e.message,
	})
}

func (e *Error) Code() int {
	if e == nil {
		return 0
	}

	return e.code
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	return e.message
}
