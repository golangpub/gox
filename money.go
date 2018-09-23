package types

import "fmt"

type Currency string

const (
	CNY Currency = "CNY"
	USD Currency = "USD"
)

type Money struct {
	Amount   int      `json:"amount"`
	Currency Currency `json:"currency"`
}

func (m *Money) String() string {
	return fmt.Sprintf("%s %d", m.Currency, m.Amount)
}
