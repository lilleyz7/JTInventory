package data

func DeleteItem(name string) error {
	db, err := connect()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM INVENTORY WHERE name = ?", name)
	if err != nil {
		return err
	}
	return nil

}
