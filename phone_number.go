package types

import "fmt"

type PhoneNumber struct {
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}

func (n *PhoneNumber) String() string {
	return fmt.Sprintf("+%s %s", n.CountryCode, n.Number)
}
