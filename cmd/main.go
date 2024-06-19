package main

import (
	"flag"
	"fmt"
	"slices"

	"github.com/lilleyz7/JTInventory/data"
	"github.com/lilleyz7/JTInventory/handlers"
)

func init() {
	err := data.Setup()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	commandList := []string{"help", "new", "add", "delete", "sub", "view"}
	command := flag.String("command", "help", "enter command\nThe following commands are available: \nhelp \nnew \nadd \nsub \ndelete \nview \nHelp can be found with -command=help")
	prodName := flag.String("name", "", "Enter the name of the product product you would like to add or interact with")
	prodQuantity := flag.Int("quantity", 0, "Quantity of the product you wish to add or interact with")

	flag.Parse()

	validCommand := slices.Contains(commandList, *command)
	if !validCommand {
		fmt.Printf("Invalid command: %s", *command)
		*command = "help"
	}

	handlers.Input(*command, *prodName, *prodQuantity)
}
