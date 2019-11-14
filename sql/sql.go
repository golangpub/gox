package sql

import (
	"database/sql"
	"github.com/gopub/log"
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
