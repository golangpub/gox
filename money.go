package types

import (
	"database/sql/driver"
	"fmt"
)

type Currency string

const (
	CNY Currency = "CNY"
	USD Currency = "USD"
)

var _ driver.Valuer = (*Money)(nil)

type Money struct {
	Amount   int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

func (m *Money) String() string {
	return fmt.Sprintf("%s %d", m.Currency, m.Amount)
}

func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("(%d,%s)", m.Amount, m.Currency), nil
}
