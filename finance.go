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
	return fmt.Sprintf("%s %s", m.Currency, m.Amount.String())
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
	return fmt.Sprintf("(%s,%s)", m.Currency, m.Amount.String()), nil
}
