package gox

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math/big"
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

	s, ok := src.(string)
	if !ok {
		var b []byte
		b, ok = src.([]byte)
		if ok {
			s = string(b)
		}
	}

	if len(s) == 0 {
		return nil
	}

	if !ok || len(s) < 10 {
		return fmt.Errorf("failed to parse %v into gox.PhoneNumber", src)
	}

	s = s[1 : len(s)-1]
	if s[len(s)-1] == ',' {
		k, _ := fmt.Sscanf(s, "%d,%d", &n.Code, &n.Number)
		if k == 2 {
			return nil
		}
	} else {
		k, _ := fmt.Sscanf(s, "%d,%d,%s", &n.Code, &n.Number, &n.Extension)
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

// Address
type Address struct {
	Country     string `json:"country,omitempty"`
	Province    string `json:"province,omitempty"`
	City        string `json:"city,omitempty"`
	District    string `json:"district,omitempty"`
	Street      string `json:"street,omitempty"`
	Building    string `json:"building,omitempty"`
	Room        string `json:"room,omitempty"`
	PostCode    string `json:"post_code,omitempty"`
	Name        string `json:"name,omitempty"`
	PostalTitle string `json:"postal_title,omitempty"`
}

func NewAddress() *Address {
	return &Address{}
}

var _ driver.Valuer = (*Address)(nil)
var _ sql.Scanner = (*Address)(nil)

func (a *Address) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok || len(b) < 2 || b[0] != '(' || b[len(b)-1] != ')' {
		return fmt.Errorf("failed to parse %v into gox.Address", src)
	}
	b = b[1 : len(b)-1]
	strs := strings.Split(string(b), ",")
	if len(strs) != 10 {
		return fmt.Errorf("failed to parse %v into gox.Address", src)
	}

	a.Country = strs[0]
	a.Province = strs[1]
	a.City = strs[2]
	a.District = strs[3]
	a.Street = strs[4]
	a.Building = strs[5]
	a.Room = strs[6]
	a.PostCode = strs[7]
	a.Name = strs[8]
	a.PostalTitle = strs[9]
	return nil
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
	name := strings.Replace(a.name, ",", "\\,", -1)
	postTitle := strings.Replace(a.PostalTitle, ",", "\\,", -1)

	s := fmt.Sprintf("(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)", country, province, city, district, street,
		building, room, postCode, name, postTitle)
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
	if len(b) < 2 || b[0] != '(' || b[len(b)-1] != ')' {
		return fmt.Errorf("parse %s into gox.Money failed", string(b))
	}
	b = b[1 : len(b)-1]
	s := strings.Replace(string(b), ",", " ", -1)
	k, err := fmt.Sscanf(s, "%s %d", &m.Currency, &m.Amount)
	if k == 2 {
		return nil
	}
	return fmt.Errorf("parse %v into gox.Money: %w", string(b), err)
}

func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%d)", m.Currency, m.Amount), nil
}

// Money
type Coin struct {
	Currency Currency `json:"currency"`
	Amount   big.Int  `json:"amount"`
}

var _ driver.Valuer = (*Coin)(nil)
var _ sql.Scanner = (*Coin)(nil)

func (c *Coin) String() string {
	return fmt.Sprintf("%s %s", c.Currency, c.Amount.String())
}

func (c *Coin) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("parse %v into gox.Coin failed", src)
	}
	if len(b) < 2 || b[0] != '(' || b[len(b)-1] != ')' {
		return fmt.Errorf("parse %s into gox.Coin failed", string(b))
	}
	b = b[1 : len(b)-1]
	s := strings.Replace(string(b), ",", " ", -1)
	var amount string
	k, err := fmt.Sscanf(s, "%s %s", &c.Currency, &amount)
	if err != nil {
		return fmt.Errorf("sscanf %s: %w", s, err)
	}
	if k != 2 {
		return fmt.Errorf("parse %v into gox.Money failed", string(b))
	}
	_, ok = c.Amount.SetString(amount, 10)
	if !ok {
		return fmt.Errorf("parse %s into big.Int failed", amount)
	}
	return nil
}

func (c *Coin) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", c.Currency, c.Amount.String()), nil
}
