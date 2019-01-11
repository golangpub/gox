package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"
)

type jsonEntityTime struct {
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type EntityTime struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *EntityTime) MarshalJSON() ([]byte, error) {
	v := &jsonEntityTime{
		CreatedAt: e.CreatedAt.Unix(),
		UpdatedAt: e.UpdatedAt.Unix(),
	}
	return json.Marshal(v)
}

func (e *EntityTime) UnmarshalJSON(data []byte) error {
	var v *jsonEntityTime
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	e.CreatedAt = time.Unix(v.CreatedAt, 0)
	e.UpdatedAt = time.Unix(v.UpdatedAt, 0)
	return nil
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
