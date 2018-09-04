package types

import "fmt"

type Currency string

const (
	CNY Currency = "CNY"
	USD Currency = "USD"
)

type Money struct {
	Cents    int      `json:"cents"`
	Currency Currency `json:"currency"`
}

func (m *Money) String() string {
	return fmt.Sprintf("%s %d", m.Currency, m.Cents)
}
