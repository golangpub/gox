package types

import "fmt"

type FullName struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

func (n *FullName) String() string {
	return fmt.Sprintf("%s %s %s", n.FirstName, n.MiddleName, n.LastName)
}
