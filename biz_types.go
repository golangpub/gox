package gox

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"

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
	CountryCode    int    `json:"country_code"`
	NationalNumber int64  `json:"national_number"`
	Extension      string `json:"extension,omitempty" sql:"type:VARCHAR(10)"`
}

var _ driver.Valuer = (*PhoneNumber)(nil)

func (n *PhoneNumber) String() string {
	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d%d", n.CountryCode, n.NationalNumber)
	}

	return fmt.Sprintf("+%d%d-%s", n.CountryCode, n.NationalNumber, n.Extension)
}

func (n *PhoneNumber) InternationalFormat() string {
	pn, err := phonenumbers.Parse(n.String(), "")
	if err != nil {
		return ""
	}
	return phonenumbers.Format(pn, phonenumbers.INTERNATIONAL)
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
		return fmt.Errorf("failed to parse %v into gox.PhoneNumber", src)
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
	return fmt.Errorf("failed to parse %s into gox.PhoneNumber", s)
}

func (n *PhoneNumber) Value() (driver.Value, error) {
	if n == nil {
		return nil, nil
	}
	ext := strings.Replace(n.Extension, ",", "\\,", -1)
	s := fmt.Sprintf("(%d,%d,%s)", n.CountryCode, n.NationalNumber, ext)
	return s, nil
}

func (n *PhoneNumber) Copy(v interface{}) error {
	if pn, ok := v.(*PhoneNumber); ok {
		*n = *pn
		return nil
	}

	if pn, ok := v.(*base.PhoneNumber); ok {
		n.CountryCode = int(pn.CountryCode)
		n.NationalNumber = pn.NationalNumber
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

// Address
type Address struct {
	Country  string `json:"country,omitempty"`
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
	District string `json:"district,omitempty"`
	Street   string `json:"street,omitempty"`
	Building string `json:"building,omitempty"`
	Room     string `json:"room,omitempty"`
	PostCode string `json:"post_code,omitempty"`

	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

var _ driver.Valuer = (*Address)(nil)
var _ sql.Scanner = (*Address)(nil)

func (a *Address) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("failed to parse %v into gox.Address", src)
	}

	k, err := fmt.Sscanf(string(b), "(%s,%s,%s,%s,%s,%s,%s,%s)", &a.Country, &a.Province, &a.City, &a.District,
		&a.Street, &a.Building, &a.Room, &a.PostCode)
	if k == 8 {
		return nil
	}
	return fmt.Errorf("failed to parse %v into gox.Address: %v", string(b), err)
}

func (a *Address) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	country := strings.Replace(a.Country, ",", "\\,", -1)
	province := strings.Replace(a.Province, ",", "\\,", -1)
	city := strings.Replace(a.City, ",", "\\,", -1)
	district := strings.Replace(a.District, ",", "\\,", -1)
	street := strings.Replace(a.Street, ",", "\\,", -1)
	building := strings.Replace(a.Building, ",", "\\,", -1)
	room := strings.Replace(a.Room, ",", "\\,", -1)
	postCode := strings.Replace(a.PostCode, ",", "\\,", -1)

	s := fmt.Sprintf("(%s,%s,%s,%s,%s,%s,%s,%s)", country, province, city, district, street, building, room, postCode)
	return s, nil
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

// Currency
type Currency string

const (
	CNY Currency = "CNY"
	USD Currency = "USD"

	ETH Currency = "ETH"
	BTC Currency = "BTC"
)

func (c Currency) Upper() Currency {
	return Currency(strings.ToUpper(string(c)))
}

// Money
type Money struct {
	Currency Currency `json:"currency"`
	Amount   int64    `json:"amount"`
}

var _ driver.Valuer = (*Money)(nil)
var _ sql.Scanner = (*Money)(nil)

func (m *Money) String() string {
	return fmt.Sprintf("%s %d", m.Currency, m.Amount)
}

func (m *Money) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("failed to parse %v into gox.Money", src)
	}

	k, err := fmt.Sscanf(string(b), "(%s,%d)", &m.Currency, &m.Amount)
	if k == 2 {
		return nil
	}
	return fmt.Errorf("failed to parse %v into gox.Money: %v", string(b), err)
}

func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%d)", m.Currency, m.Amount), nil
}
