package gox

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math/big"
	"strings"
)

// Money
type Money struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
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
	Currency string  `json:"currency"`
	Amount   big.Int `json:"amount"`
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
