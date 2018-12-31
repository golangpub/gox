package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type Currency string

const (
	CNY Currency = "CNY"
	USD Currency = "USD"
)

var _ driver.Valuer = (*Money)(nil)

type Money struct {
	Currency Currency `json:"currency"`
	Amount   int64    `json:"amount"`
}

func (m *Money) String() string {
	return fmt.Sprintf("%s %d", m.Currency, m.Amount)
}

func (m *Money) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("failed to parse %v into types.Money", src))
	}

	_, err := fmt.Sscanf(string(b), "(%s,%d)", &m.Currency, &m.Amount)
	return err
}

func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%d)", m.Currency, m.Amount), nil
}
