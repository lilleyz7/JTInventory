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
		err := data.UpdateField(name, quantity)
		if err != nil {
			fmt.Printf("Error adding to %s, with error :%s", name, err.Error())
		} else {
			fmt.Printf("Successfully added %d to %s", quantity, name)
		}
	case "sub":
		q := quantity * -1
		err := data.UpdateField(name, q)
		if err != nil {
			fmt.Printf("Error removing from %s, with error :%s", name, err.Error())
		} else {
			fmt.Printf("Successfully removed %d from %s", q, name)
		}
	case "delete":

	default:
		fmt.Printf("The command '%s' is not supported.\nPlease enter add, sub, delete, or new", command)
	}
}
