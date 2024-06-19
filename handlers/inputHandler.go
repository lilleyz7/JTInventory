package handlers

import (
	"fmt"

	"github.com/lilleyz7/JTInventory/data"
)

func Input(command string, name string, quantity int) {
	switch command {
	case "new":
		err := data.CreateField(name, quantity)
		if err != nil {
			fmt.Printf("Error creating %s, with error :%s", name, err.Error())
		} else {
			fmt.Printf("Successfully created %s", name)
		}

	case "add":
		err := data.UpdateQuantity(name, quantity)
		if err != nil {
			fmt.Printf("Error adding to %s, with error :%s", name, err.Error())
		} else {
			fmt.Printf("Successfully added %d to %s", quantity, name)
		}
	case "sub":
		q := quantity * -1
		err := data.UpdateQuantity(name, q)
		if err != nil {
			fmt.Printf("Error removing from %s, with error :%s", name, err.Error())
		} else {
			fmt.Printf("Successfully removed %d from %s", q, name)
		}
	case "delete":
		err := data.DeleteItem(name)
		if err != nil {
			fmt.Printf("Failed to delete item: %s, with error: %s", name, err.Error())
		} else {
			fmt.Printf("Successfully product: %s", name)
		}

	case "view":
		data.ViewData(name)

	default:
		fmt.Printf("The command '%s' is not supported.\nPlease enter add, sub, delete, or new", command)
	}
}
