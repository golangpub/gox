package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

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

var _ driver.Valuer = (*Money)(nil)
var _ sql.Scanner = (*Money)(nil)

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

	k, err := fmt.Sscanf(string(b), "(%s,%d)", &m.Currency, &m.Amount)
	if k == 2 {
		return nil
	}
	return errors.New(fmt.Sprintf("failed to parse %v into types.Money: %v", string(b), err))
}

func (m *Money) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%d)", m.Currency, m.Amount), nil
}
