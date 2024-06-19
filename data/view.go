package data

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/lilleyz7/JTInventory/models"
)

func ViewData(name string) {
	if name == "" {
		viewAllProduct()
	} else {
		getSingleProduct(name)
	}
}

func getSingleProduct(name string) {
	var product models.Product

	db, err := connect()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM inventory WHERE product_name = ?")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	row := stmt.QueryRow(name)

	err = row.Scan(&product.Id, &product.Name, &product.Quantity, &product.Created_at, &product.Updated_at)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	createTable([]models.Product{product})

}

func viewAllProduct() {
	var products []models.Product

	db, err := connect()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM inventory")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Quantity, &product.Created_at, &product.Updated_at)
		fmt.Printf("Error: %v\n", err)
		products = append(products, product)
	}

	createTable(products)
}

func createTable(products []models.Product) {
	columns := []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Name", Width: 15},
		{Title: "Quantity", Width: 15},
		{Title: "Created_at", Width: 20},
		{Title: "Updated_at", Width: 20},
	}
	var rows = []table.Row{}

	for _, product := range products {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", product.Id),
			product.Name,
			fmt.Sprintf("%d", product.Quantity),
			product.Created_at.Format("2006-01-02T15:04:05"),
			product.Updated_at.Format("2006-01-02T15:04:05"),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(len(products)),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("200")).BorderBottom(true)
	t.SetStyles(s)
	fmt.Println("\n\n\n", t.View(), "\n\n\n")
}
