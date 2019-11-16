package gox

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gopub/gox/sql"

	"github.com/gopub/gox/protobuf/base"
	"github.com/nyaruka/phonenumbers"
)

type FullName struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

var _ driver.Valuer = (*FullName)(nil)

func (n *FullName) String() string {
	return fmt.Sprintf("%s %s %s", n.FirstName, n.MiddleName, n.LastName)
}

func (n *FullName) Scan(src interface{}) error {
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

	if !ok || len(s) < 4 {
		return fmt.Errorf("failed to parse %v into gox.PhoneNumber", src)
	}

	s = s[1 : len(s)-1]
	segments := strings.Split(s, ",")
	if len(segments) != 3 {
		return fmt.Errorf("failed to parse %v into gox.PhoneNumber", src)
	}

	n.FirstName, n.MiddleName, n.LastName = segments[0], segments[1], segments[2]
	return nil
}

func (n *FullName) Value() (driver.Value, error) {
	if n == nil {
		return nil, nil
	}
	s := fmt.Sprintf("(%s,%s,%s)", n.FirstName, n.MiddleName, n.LastName)
	return s, nil
}

// PhoneNumber
type PhoneNumber struct {
	Code      int    `json:"code"`
	Number    int64  `json:"number"`
	Extension string `json:"extension,omitempty" sql:"type:VARCHAR(10)"`
}

var _ driver.Valuer = (*PhoneNumber)(nil)

func (n *PhoneNumber) String() string {
	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d%d", n.Code, n.Number)
	}

	return fmt.Sprintf("+%d%d-%s", n.Code, n.Number, n.Extension)
}

func (n *PhoneNumber) InternationalFormat() string {
	pn, err := phonenumbers.Parse(n.String(), "")
	if err != nil {
		return ""
	}
	return phonenumbers.Format(pn, phonenumbers.INTERNATIONAL)
}

func (n *PhoneNumber) MaskString() string {
	nnBytes := []byte(fmt.Sprint(n.Number))
	maskLen := (len(nnBytes) + 2) / 3
	start := len(nnBytes) - 2*maskLen
	for i := 0; i < maskLen; i++ {
		nnBytes[start+i] = '*'
	}

	nn := string(nnBytes)

	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d%s", n.Code, nn)
	}

	return fmt.Sprintf("+%d%s-%s", n.Code, nn, n.Extension)
}

func (n *PhoneNumber) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	s, err := ParseString(src)
	if err != nil {
		return fmt.Errorf("parse string: %w", err)
	}
	if len(s) == 0 {
		return nil
	}

	fields, err := sql.ParseCompositeFields(s)
	if err != nil {
		return fmt.Errorf("parse composite fields %s: %w", s, err)
	}

	if len(fields) != 3 {
		return fmt.Errorf("parse composite fields %s: got %v", s, fields)
	}

	code, err := ParseInt(fields[0])
	if err != nil {
		return fmt.Errorf("parse code %s: %w", fields[0], err)
	}
	n.Code = int(code)
	n.Number, err = ParseInt(fields[1])
	if err != nil {
		return fmt.Errorf("parse code %s: %w", fields[1], err)
	}
	n.Extension = fields[2]
	return nil
}

func (n *PhoneNumber) Value() (driver.Value, error) {
	if n == nil {
		return nil, nil
	}
	ext := strings.Replace(n.Extension, ",", "\\,", -1)
	s := fmt.Sprintf("(%d,%d,%s)", n.Code, n.Number, ext)
	return s, nil
}

func (n *PhoneNumber) Copy(v interface{}) error {
	if pn, ok := v.(*PhoneNumber); ok {
		*n = *pn
		return nil
	}

	if pn, ok := v.(*base.PhoneNumber); ok {
		n.Code = int(pn.CountryCode)
		n.Number = pn.NationalNumber
		n.Extension = pn.Extension
		return nil
	}

	var s string
	if b, ok := v.([]byte); ok {
		s = string(b)
	} else if s, ok = v.(string); !ok {
		return fmt.Errorf("v is %v instead of string or []byte", reflect.TypeOf(v))
	}

	res, err := ParsePhoneNumber(s)
	if err != nil {
		return err
	}
	*n = *res
	return nil
}

func NewPhoneNumber(callingCode int, number int64) *PhoneNumber {
	pn := new(PhoneNumber)
	pn.Code = callingCode
	pn.Number = number
	return pn
}

func ParsePhoneNumber(s string) (*PhoneNumber, error) {
	parsedNumber, err := phonenumbers.Parse(s, "")
	if err != nil {
		return nil, err
	}

	if phonenumbers.IsValidNumber(parsedNumber) {
		return &PhoneNumber{
			Code:      int(parsedNumber.GetCountryCode()),
			Number:    int64(parsedNumber.GetNationalNumber()),
			Extension: parsedNumber.GetExtension(),
		}, nil
	}

	return nil, errors.New("invalid phone number")
}

// Gender
type Gender int

const (
	Male   Gender = 1
	Female Gender = 2
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}

func (g Gender) IsValid() bool {
	switch g {
	case Male, Female:
		return true
	default:
		return false
	}
}
