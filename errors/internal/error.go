package internal

import (
	"encoding/json"
	"fmt"
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

func (a *Error) UnmarshalJSON(b []byte) error {
	var obj *errorJSONObject
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	a.code = obj.Code
	a.message = obj.Message
	return nil
}

func (a Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&errorJSONObject{
		Code:    a.code,
		Message: a.message,
	})
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
	return fmt.Sprintf("%d:%s", e.code, e.message)
}
