package sql

import (
	"database/sql"
	"github.com/gopub/log"
)

type (
	DB          = sql.DB
	Tx          = sql.Tx
	TxOptions   = sql.TxOptions
	Stmt        = sql.Stmt
	Row         = sql.Row
	Rows        = sql.Rows
	Conn        = sql.Conn
	Result      = sql.Result
	NullInt64   = sql.NullInt64
	NullTime    = sql.NullTime
	NullBool    = sql.NullBool
	NullFloat64 = sql.NullFloat64
	NullInt32   = sql.NullInt32
	NullString  = sql.NullString
)

var (
	ErrNoRows   = sql.ErrNoRows
	ErrTxDone   = sql.ErrTxDone
	ErrConnDone = sql.ErrConnDone
)

type ColumnScanner interface {
	Scan(dest ...interface{}) error
}

type Executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func OpenPostgres(dbURL string) *sql.DB {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Open %s: %+v", dbURL, err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping %s: %+v", dbURL, err)
	}
	return db
}
