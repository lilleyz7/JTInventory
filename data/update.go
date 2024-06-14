package data

func UpdateQuantity(name string, value int) error {
	db, err := connect()
	if err != nil {
		return err
	}

	defer db.Close()
	
	stmt, err := db.Prepare("UPDATE inventory SET product_quantity = product_quantity + ? WHERE name = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the change value and product ID
	_, err = stmt.Exec(value, name)
	if err != nil {
		return err
	}

	return nil
}
