package models

type Product struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Media     string `json:"media" db:"media"`
	Price     int    `json:"price" db:"price"`
	BulkPrice int    `json:"bulk_price" db:"bulk_price"`
	InStock   int    `json:"in_stock" db:"in_stock"`
}
