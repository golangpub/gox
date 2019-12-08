package gox

import (
	"database/sql/driver"
	"fmt"

	"github.com/gopub/gox/sql"
	"github.com/shopspring/decimal"
)

// Money
type Money struct {
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
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

	if len(fields) != 2 {
		return fmt.Errorf("parse composite fields %s: got %v", s, fields)
	}
	m.Currency = fields[0]
	m.Amount, err = decimal.NewFromString(fields[1])
	if err != nil {
		return fmt.Errorf("parse amount %s: %w", fields[1], err)
	}
	return nil
}

func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%d)", m.Currency, m.Amount), nil
}

// Money
type Coin struct {
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
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

	if len(fields) != 2 {
		return fmt.Errorf("parse composite fields %s: got %v", s, fields)
	}
	c.Currency = fields[0]
	c.Amount, err = decimal.NewFromString(fields[1])
	if err != nil {
		return fmt.Errorf("parse amount %s: %w", fields[1], err)
	}
	return nil
}

func (c *Coin) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", c.Currency, c.Amount.String()), nil
}
