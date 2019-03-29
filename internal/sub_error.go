package internal

import (
	"encoding/json"
	"fmt"
)

type subErrorJSONObject struct {
	Code    int    `json:"code"`
	SubCode int    `json:"sub_code"`
	Message string `json:"message"`
}

// SubError represents error info
type SubError struct {
	code    int
	subCode int
	message string
}

func NewSubError(code, subCode int, msg string) *SubError {
	return &SubError{code: code, subCode: subCode, message: msg}
}

func (e *SubError) UnmarshalJSON(b []byte) error {
	var obj *subErrorJSONObject
	if err := json.Unmarshal(b, &obj); err != nil {
		return err
	}

	e.code = obj.Code
	e.subCode = obj.SubCode
	e.message = obj.Message
	return nil
}

func (e *SubError) MarshalJSON() ([]byte, error) {
	return json.Marshal(&subErrorJSONObject{
		Code:    e.code,
		SubCode: e.subCode,
		Message: e.message,
	})
}

func (e *SubError) Code() int {
	if e == nil {
		return 0
	}

	return e.code
}

func (e *SubError) SubCode() int {
	if e == nil {
		return 0
	}

	return e.subCode
}

func (e *SubError) Message() string {
	if e == nil {
		return ""
	}

	return e.message
}

func (e *SubError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%d:%d:%s", e.code, e.subCode, e.message)
}
