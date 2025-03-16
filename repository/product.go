package repository

import (
	"app/models"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(
	db *sqlx.DB,
) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Insert(product *models.Product) error {
	const query = `
    insert into products (
      id, name, media, price,
      bulk_price, in_stock
    )
    values (
      :id, :name, :media, :price,
      :bulk_price, :in_stock
    )
  `
	_, err := r.db.NamedExec(query, product)
	if err != nil {
		return fmt.Errorf("Error while inserting product: %w", err)
	}
	return nil
}

func (r *ProductRepo) GetAll() ([]models.Product, error) {
	const query = `
    select * from products
  `
	products := make([]models.Product, 0)
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting all products: %w", err)
	}
	return products, nil
}

func (r *ProductRepo) GetByID(id string) (*models.Product, error) {
	const query = `
    select * from products where id=$1
  `
	product := &models.Product{}
	err := r.db.Get(product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting product by ID: %w", err)
	}
	return product, nil
}

// to be implemented
func (r *ProductRepo) Update(id string, product *models.Product) error {
	return nil
}
