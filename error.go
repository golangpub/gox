package types

import (
	"strconv"
	"strings"
)

type Error interface {
	error
	Code() Ecode
	Msg() string
	ToMap() M
}

// errorInfo represents error info
type errorInfo struct {
	code Ecode
	msg  string
}

func (e *errorInfo) Code() Ecode {
	return e.code
}

func (e *errorInfo) Msg() string {
	if e == nil {
		return ""
	}
	return e.msg
}

func (e *errorInfo) Error() string {
	if e == nil {
		return ""
	}
	return e.code.String() + ":" + e.msg
}

func (e *errorInfo) ToMap() M {
	return M{"code": e.code, "msg": e.msg}
}

func NewError(code Ecode, msg string) Error {
	return &errorInfo{code: code, msg: msg}
}

func ParseError(err error) Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(Error); ok {
		return e
	}

	strs := strings.SplitN(err.Error(), ":", 2)
	if len(strs) < 2 {
		return ErrUnknown
	}

	if i, e := strconv.Atoi(strs[0]); e == nil {
		switch Ecode(i) {
		case EcodeSuccess:
			if strs[1] == EcodeSuccess.String() {
				return ErrSuccess
			}
		case EcodeParam:
			if strs[1] != EcodeParam.String() || len(strs) < 3 {
				break
			}
			return ErrParam(strs[2])
		case EcodeAuth:
			if strs[1] == ErrAuth.msg {
				return ErrAuth
			}
		case EcodeForbidden:
			if strs[1] != EcodeForbidden.String() || len(strs) < 3 {
				break
			}
			return ErrForbidden(strs[2])
		case EcodeDupKey:
			if strs[1] != EcodeDupKey.String() || len(strs) < 3 {
				break
			}
			return ErrDupKey(strs[2])
		case EcodeServer:
			if strs[1] == ErrServer.msg {
				return ErrServer
			}
		case EcodeNotFound:
			if strs[1] != EcodeNotFound.String() || len(strs) < 3 {
				break
			}
			return ErrNotFound(strs[2])
		case EcodeExpired:
			if strs[1] != EcodeExpired.String() {
				break
			}
			return ErrExpired(strs[2])
		}
	}
	return ErrUnknown
}

type Ecode int

const (
	// EcodeSuccess represents no error
	EcodeSuccess Ecode = 0

	// EcodeUnknown represents unknown error
	EcodeUnknown Ecode = 1

	// EcodeServer represents internal server error
	EcodeServer Ecode = 1001

	// EcodeBadRequest represents invalid app id/key, request sign, etc.
	EcodeBadRequest Ecode = 2001

	// EcodeParam represents lack of parameter or invalid value
	EcodeParam Ecode = 2002

	// EcodeDupKey represents duplicate key
	EcodeDupKey Ecode = 2003

	// EcodeExpired represents expiration such as captcha, invitation link, etc.
	EcodeExpired Ecode = 2005

	// EcodeNotFound represents data not found
	EcodeNotFound Ecode = 2006

	// EcodeAuth represents session expired or login failed
	EcodeAuth Ecode = 2101

	// EcodeForbidden represents operation is forbidden
	EcodeForbidden Ecode = 2102
)

func (e Ecode) String() string {
	switch e {
	case EcodeSuccess:
		return "success"
	case EcodeServer:
		return "server"
	case EcodeBadRequest:
		return "bad_request"
	case EcodeParam:
		return "invalid_param"
	case EcodeDupKey:
		return "duplicate"
	case EcodeExpired:
		return "expired"
	case EcodeNotFound:
		return "not_found"
	case EcodeAuth:
		return "auth"
	case EcodeForbidden:
		return "forbidden"
	default:
		return "unknown"
	}
}

func (e Ecode) Int() int {
	return int(e)
}

var ErrSuccess = &errorInfo{code: EcodeSuccess, msg: EcodeSuccess.String()}
var ErrUnknown = &errorInfo{code: EcodeUnknown, msg: EcodeUnknown.String()}
var ErrAuth = &errorInfo{code: EcodeAuth, msg: EcodeAuth.String()}
var ErrServer = &errorInfo{code: EcodeServer, msg: EcodeServer.String()}

func ErrParam(param string) Error {
	return &errorInfo{code: EcodeParam, msg: param}
}

func ErrDupKey(key string) Error {
	return &errorInfo{code: EcodeDupKey, msg: key}
}

func ErrForbidden(msg string) Error {
	return &errorInfo{code: EcodeForbidden, msg: msg}
}

func ErrExpired(msg string) Error {
	return &errorInfo{code: EcodeForbidden, msg: msg}
}

func ErrBadRequest(msg string) Error {
	return &errorInfo{code: EcodeBadRequest, msg: msg}
}

func ErrNotFound(msg string) Error {
	return &errorInfo{code: EcodeNotFound, msg: msg}
}
