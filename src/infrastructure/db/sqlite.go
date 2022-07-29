package db

import (
	interfaces "audioTest/src/interfaces/repositories"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type SqliteHandler struct {
	Conn *sql.DB
}

func (handler *SqliteHandler) Execute(statement string) {
	result, err := handler.Conn.Exec(statement)
	_, _ = result, err
}

func (handler *SqliteHandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow)
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}

func NewSqliteHandler(dbfileName string) *SqliteHandler {
	conn, err := sql.Open("sqlite3", dbfileName)
	if err != nil {
		// Fail fast
		log.Fatal(err)
	}

	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}
