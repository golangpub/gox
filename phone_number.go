package types

import "fmt"

type PhoneNumber struct {
	CountryCode    int    `json:"country_code"`
	NationalNumber int64  `json:"national_number"`
	Extension      string `json:"extension,omitempty"`
}

func (n *PhoneNumber) String() string {
	if len(n.Extension) == 0 {
		return fmt.Sprintf("+%d-%d", n.CountryCode, n.NationalNumber)
	}

	return fmt.Sprintf("+%d-%d-%s", n.CountryCode, n.NationalNumber, n.Extension)
}
