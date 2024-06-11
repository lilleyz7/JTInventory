package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lilleyz7/JTInventory/data"
)

func main() {
	err := data.Setup()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	prodName := flag.String("name", "", "enter the name of the product")
	prodQuantity := flag.Int("quantity", 0, "quantity of the product")

	command := os.Args[1]
}
