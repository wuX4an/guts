package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "database.sql")
	if err != nil {
		return nil, err
	}
	return db, nil
}
