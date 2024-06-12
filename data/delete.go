package data

import "database/sql"

func DeleteItem(name string) error {
	db, err := sql.Open("sqlite", "storage.db")
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM INVENTORY WHERE name = ?", name)
	if err != nil {
		return err
	}
	return nil

}
