package database

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema/schema.sql
var schema string

func GetInstance() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "file:database.sqlite?cache=shared")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
