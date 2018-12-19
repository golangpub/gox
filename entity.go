package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"time"
)

type BaseEntity struct {
	ID        ID        `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SQLBigInt big.Int

func (i *SQLBigInt) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	var s string
	var ok bool
	s, ok = src.(string)
	if !ok {
		b, ok := src.([]byte)
		if ok {
			s = string(b)
		}
	}

	if !ok {
		return errors.New(fmt.Sprintf("failed to parse %v into big.Int", src))
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
