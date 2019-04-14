package gox

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
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
		return fmt.Errorf("failed to parse %v into big.Int", src)
	}

	_, ok = (*big.Int)(i).SetString(s, 10)
	if !ok {
		return fmt.Errorf("failed to parse %v into big.Int", src)
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

type SQLExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type SQLMap map[string]interface{}

var _ sql.Scanner = (*SQLMap)(nil)
var _ driver.Valuer = SQLMap(nil)

func (m *SQLMap) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		s, ok := src.(string)
		if ok {
			b = []byte(s)
		}
	}

	if !ok {
		return fmt.Errorf("failed to parse %v into gox.SQLMap", src)
	}

	return JSONUnmarshal(b, m)
}

func (m SQLMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

func MustPrepareStmt(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Panic(err)
	}
	return stmt
}
