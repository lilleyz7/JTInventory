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
	command := flag.String("command", "help", "enter command \n commands can be found with -command=help")
	prodName := flag.String("name", "", "enter the name of the product")
	prodQuantity := flag.Int("quantity", 0, "quantity of the product")

	flag.Parse()
	if *command == "help" {
		printHelp()
	}

	validCommand := slices.Contains(commandList, *command)
	if !validCommand {
		printHelp()
	}

	handlers.Input(*command, *prodName, *prodQuantity)
}

func printHelp() {

}
