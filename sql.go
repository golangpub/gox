package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"math/big"
	"time"
)

type SQLRowTime struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (e *SQLRowTime) EntityTime() *EntityTime {
	return &EntityTime{
		CreatedAt: e.CreatedAt.Unix(),
		UpdatedAt: e.UpdatedAt.Unix(),
	}
}

type SQLBaseRow struct {
	ID      ID   `json:"-"`
	Deleted bool `json:"-"`
	SQLRowTime
}

func (e *SQLBaseRow) Entity() *BaseEntity {
	return &BaseEntity{
		ID:         e.ID,
		EntityTime: *e.SQLRowTime.EntityTime(),
	}
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

func (i *SQLBigInt) Int() *big.Int {
	return (*big.Int)(i)
}

type SQLRowScanner interface {
	Scan(dest ...interface{}) error
}
