package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

var _ driver.Valuer = (*PhoneNumber)(nil)

type PhoneNumber struct {
	CountryCode    int    `json:"country_code"`
	NationalNumber int64  `json:"national_number"`
	Extension      string `json:"extension,omitempty" sql:"type:VARCHAR(10)"`
}

func (n *PhoneNumber) String() string {
	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d%d", n.CountryCode, n.NationalNumber)
	}

	return fmt.Sprintf("+%d%d-%s", n.CountryCode, n.NationalNumber, n.Extension)
}

func (n *PhoneNumber) MaskString() string {
	nnBytes := []byte(fmt.Sprint(n.NationalNumber))
	maskLen := (len(nnBytes) + 2) / 3
	start := len(nnBytes) - 2*maskLen
	for i := 0; i < maskLen; i++ {
		nnBytes[start+i] = '*'
	}

	nn := string(nnBytes)

	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d%s", n.CountryCode, nn)
	}

	return fmt.Sprintf("+%d%s-%s", n.CountryCode, nn, n.Extension)
}

func (n *PhoneNumber) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	s, ok := src.(string)
	if !ok {
		var b []byte
		b, ok = src.([]byte)
		if ok {
			s = string(b)
		}
	}

	if !ok || len(s) < 10 {
		return errors.New(fmt.Sprintf("failed to parse %v into types.PhoneNumber", src))
	}

	s = s[1 : len(s)-1]
	if s[len(s)-1] == ',' {
		k, _ := fmt.Sscanf(s, "%d,%d", &n.CountryCode, &n.NationalNumber)
		if k == 2 {
			return nil
		}
	} else {
		k, _ := fmt.Sscanf(s, "%d,%d,%s", &n.CountryCode, &n.NationalNumber, &n.Extension)
		if k == 3 {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("failed to parse %s into types.PhoneNumber", s))
}

func (n *PhoneNumber) Value() (driver.Value, error) {
	if n == nil {
		return nil, nil
	}
	ext := strings.Replace(n.Extension, ",", "\\,", -1)
	s := fmt.Sprintf("(%d,%d,%s)", n.CountryCode, n.NationalNumber, ext)
	return s, nil
}
