package data

func CreateField(name string, startingQuant int) error {
	db, err := connect()
	if err != nil {
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO inventory(product_name, product_quantity)
							VALUES(?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, startingQuant)
	if err != nil {
		return err
	}

	return nil

}
