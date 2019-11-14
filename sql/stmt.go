package sql

import (
	"database/sql"
	"fmt"
	"log"
)

func MustPrepare(db *sql.DB, format string, args ...interface{}) *sql.Stmt {
	stmt, err := db.Prepare(fmt.Sprintf(format, args...))
	if err != nil {
		log.Fatalf("Prepare: %+v", err)
	}
	return stmt
}
