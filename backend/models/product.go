package models

type Product struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Price       int    `json:"price" db:"price"`
	BulkPrice   int    `json:"bulk_price" db:"bulk_price"`
	IsAvailable bool   `json:"is_available" db:"is_available"`
}
