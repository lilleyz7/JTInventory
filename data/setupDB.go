package data

import (
	"database/sql"
)

func Setup() error {
	db, err := sql.Open("sqlite", "storage.db")
	if err != nil {
		return err
	}
	defer db.Close()

	tableSQL := `CREATE TABLE IF NOT EXISTS inventory (
				product_id INTEGER PRIMARY KEY,
				product_name TEXT NOT NULL UNIQUE,
				product_quantity INT DEFAULT 0,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				last_updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
				)`

	_, err = db.Exec(tableSQL)
	if err != nil {
		return err
	}

	return nil
}
