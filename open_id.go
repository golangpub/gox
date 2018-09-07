package types

import "fmt"

type OpenID struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
}

func (o *OpenID) String() string {
	return fmt.Sprintf("%s_%s", o.Provider, o.ID)
}
