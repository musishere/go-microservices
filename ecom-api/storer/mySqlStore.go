package storer

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type MySQLStorer struct {
	db *sqlx.DB
}

func NewMySQLStorer(db *sqlx.DB) *MySQLStorer {
	return &MySQLStorer{db: db}
}

func (d *MySQLStorer) CreateProduct(ctx context.Context, product *Product) (*Product, error) {
	res, err := d.db.NamedExecContext(ctx, `
		INSERT INTO products (name, image, category, description, rating, num_reviews, price, count_in_stock)
		VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)
		`, product)

	if err != nil {
		log.Println("Error creating product", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id", err)
		return nil, err
	}

	product.ID = id

	return product, nil
}

func (d *MySQLStorer) GetProductByID(ctx context.Context, id int64) (*Product, error) {
	var product Product
	err := d.db.GetContext(ctx, &product, `
		SELECT * FROM products WHERE id = ?
		`, id)
	if err != nil {
		log.Println("Error getting product", err)
		return nil, err
	}

	return &product, nil
}

func (d *MySQLStorer) ListProducts(ctx context.Context) ([]*Product, error) {
	var products []*Product
	err := d.db.SelectContext(ctx, &products, `
		SELECT * FROM products
		`)
	if err != nil {
		log.Println("Error getting products", err)
		return nil, err
	}
	return products, nil
}

func (d *MySQLStorer) UpdateProduct(ctx context.Context, product *Product) (*Product, error) {
	_, err := d.db.NamedExecContext(ctx, `
		UPDATE products SET name = :name, image = :image, category = :category, description = :description, rating = :rating, num_reviews = :num_reviews, price = :price, count_in_stock = :count_in_stock WHERE id = :id
		`, product)
	if err != nil {
		log.Println("Error updating product", err)
		return nil, err
	}
	return product, nil
}

func (d *MySQLStorer) DeleteProduct(ctx context.Context, id int64) error {
	_, err := d.db.ExecContext(ctx, `
		DELETE FROM products WHERE id = ?
		`, id)
	if err != nil {
		log.Println("Error deleting product", err)
		return err
	}
	return nil
}
