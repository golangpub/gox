package gox

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func ParseBool(i interface{}) (bool, error) {
	if i == nil {
		return false, ErrNotExist
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), nil
	case reflect.String:
		str := strings.ToLower(v.String())
		if str == "true" {
			return true, nil
		}
		if str == "false" {
			return false, nil
		}
		return false, ErrNotExist
	default:
		b, err := ParseInt(i)
		if err == nil {
			return b != 0, nil
		}

		if v.Kind() == reflect.String {
			if str, ok := i.(string); ok {
				if str == "true" {
					return true, nil
				}

				if str == "false" {
					return false, nil
				}
			}
		}

		return false, err
	}
}

func ParseInt(i interface{}) (int64, error) {
	if i == nil {
		return 0, ErrNotExist
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r := int64(v.Uint())
		return r, nil
	case reflect.Float32, reflect.Float64:
		return int64(v.Float()), nil
	case reflect.String:
		if num, ok := i.(json.Number); ok {
			n, e := num.Int64()
			if e != nil {
				var f float64
				f, e = num.Float64()
				n = int64(f)
			}
			return n, e
		}

		if n, err := strconv.ParseInt(v.String(), 0, 64); err == nil {
			return n, nil
		}

		if n, err := strconv.ParseFloat(v.String(), 64); err == nil {
			return int64(n), nil
		}
	}
	return 0, ErrNotExist
}

func ParseFloat(i interface{}) (float64, error) {
	if i == nil {
		return 0, ErrNotExist
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	case reflect.String:
		if num, ok := i.(json.Number); ok {
			return num.Float64()
		}
		return strconv.ParseFloat(v.String(), 64)
	default:
		return 0, ErrNotExist
	}
}
