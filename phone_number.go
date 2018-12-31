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
		return fmt.Sprintf("+%d-%d", n.CountryCode, n.NationalNumber)
	}

	return fmt.Sprintf("+%d-%d-%s", n.CountryCode, n.NationalNumber, n.Extension)
}

func (n *PhoneNumber) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("failed to parse %v into types.PhoneNumber", src))
	}

	_, err := fmt.Sscanf(string(b), "(%d,%d,%s)", &n.CountryCode, &n.NationalNumber, &n.Extension)
	return err
}

func (n *PhoneNumber) Value() (driver.Value, error) {
	if n == nil {
		return nil, nil
	}
	ext := strings.Replace(n.Extension, ",", "\\,", -1)
	s := fmt.Sprintf("(%d,%d,%s)", n.CountryCode, n.NationalNumber, ext)
	return s, nil
}
