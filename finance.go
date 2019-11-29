package gox

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/gopub/gox/sql"
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
	m.Amount, err = ParseInt(fields[1])
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
	_, ok := c.Amount.SetString(fields[1], 10)
	if !ok {
		return fmt.Errorf("parse amount %s failed", fields[1])
	}
	return nil
}

func (c *Coin) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", c.Currency, c.Amount.String()), nil
}

func (c *Coin) UnmarshalJSON(data []byte) error {
	type CoinType Coin
	var cc *CoinType
	if err := json.Unmarshal(data, &cc); err == nil {
		c.Currency = cc.Currency
		c.Amount = cc.Amount
		return nil
	}

	var v struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	}
	err := json.Unmarshal(data, &v)
	if err == nil {
		c.Currency = v.Currency
		_, ok := c.Amount.SetString(v.Amount, 10)
		if !ok {
			return fmt.Errorf("cannot parse %s into big.Int", v.Amount)
		}
		return nil
	}
	return err
}
