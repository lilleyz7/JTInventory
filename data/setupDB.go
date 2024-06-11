package data

import (
	"database/sql"
)

func Setup() error {
	_, err := sql.Open("sqlite", "storage.db")
	if err != nil {
		return err
	}

	return nil
}
