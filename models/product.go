package models

import "time"

type Product struct {
	Id         int       `json:"product_id"`
	Name       string    `json:"product_name"`
	Quantity   int       `json:"product_quantity"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"last_updated"`
}
