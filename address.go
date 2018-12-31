package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

var _ driver.Valuer = (*PhoneNumber)(nil)
var _ sql.Scanner = (*PhoneNumber)(nil)

type Address struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Street   string `json:"street"`
	Building string `json:"building"`
	Room     string `json:"room"`
	PostCode string `json:"post_code"`
}

func (a *Address) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("failed to parse %v into types.Address", src))
	}

	_, err := fmt.Sscanf(string(b), "(%s,%s,%s,%s,%s,%s,%s,%s)", &a.Country, &a.Province, &a.City, &a.District,
		&a.Street, &a.Building, &a.Room, &a.PostCode)
	return err
}

func (a *Address) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}

	country := strings.Replace(a.Country, ",", "\\,", -1)
	province := strings.Replace(a.Province, ",", "\\,", -1)
	city := strings.Replace(a.City, ",", "\\,", -1)
	district := strings.Replace(a.District, ",", "\\,", -1)
	street := strings.Replace(a.Street, ",", "\\,", -1)
	building := strings.Replace(a.Building, ",", "\\,", -1)
	room := strings.Replace(a.Room, ",", "\\,", -1)
	postCode := strings.Replace(a.PostCode, ",", "\\,", -1)

	s := fmt.Sprintf("(%s,%s,%s,%s,%s,%s,%s,%s)", country, province, city, district, street, building, room, postCode)
	return s, nil
}
