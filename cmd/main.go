package main

import (
	"flag"
	"fmt"

	"github.com/lilleyz7/JTInventory/data"
)

func init() {
	err := data.Setup()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {

	// prodName := flag.String("name", "", "enter the name of the product")
	// prodQuantity := flag.Int("quantity", 0, "quantity of the product")

	// command := os.Args[1]

	flag.Parse()
}
