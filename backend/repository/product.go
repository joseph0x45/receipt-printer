package repository

import (
	"backend/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) Insert(product *models.Product) error {
	const query = `
    insert into products (
      id, name, price, bulk_price, is_available
    )
    values (
      :id, :name, :price, :bulk_price, :is_available
    )
  `
	_, err := r.db.NamedExec(query, product)
	if err != nil {
		return fmt.Errorf("Error while inserting product: %w", err)
	}
	return nil
}

func (r *ProductRepo) GetAll() ([]models.Product, error) {
	const query = "select * from products"
	data := make([]models.Product, 0)
	err := r.db.Select(&data, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all products: %w", err)
	}
	return data, nil
}
