package data

import (
	"database/sql"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "storage.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}
