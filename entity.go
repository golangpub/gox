package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"time"
)

type EntityTime struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BaseEntity struct {
	EntityTime
	ID      ID   `json:"id"`
	Deleted bool `json:"-"`
}

type SQLBigInt big.Int

var _ driver.Valuer = (*SQLBigInt)(nil)
var _ sql.Scanner = (*SQLBigInt)(nil)

func (i *SQLBigInt) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var s string
	var ok bool
	s, ok = src.(string)
	if !ok {
		var b []byte
		b, ok = src.([]byte)
		if ok {
			s = string(b)
		}
	}

	if !ok {
		return errors.New(fmt.Sprintf("failed to parse %v into big.Int", src))
	}

	if s == "null" {
		return nil
	}

	_, ok = (*big.Int)(i).SetString(s, 10)
	if !ok {
		return errors.New(fmt.Sprintf("failed to parse %v into big.Int", src))
	}
	return nil
}

func (i *SQLBigInt) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return (*big.Int)(i).String(), nil
}
