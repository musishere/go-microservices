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

func (d *MySQLStorer) CreateOrder(ctx context.Context, order *Order) (*Order, error) {

	err := d.execTx(ctx, func(tx *sqlx.Tx) error {
		order, err := d.insertOrder(ctx, tx, order)
		if err != nil {
			log.Println("Error creating an order", err)
			return err
		}

		for _, item := range order.Items {
			item.OrderID = order.ID
			_, err := d.insertOrderItem(ctx, tx, item)
			if err != nil {
				log.Println("Error creating an order item", err)
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Println("Error creating an order", err)
		return nil, err
	}

	return order, nil

}

func (d *MySQLStorer) GetOrderByID(ctx context.Context, id int64) (*Order, error) {
	var order Order
	err := d.db.GetContext(ctx, &order, `
		SELECT * FROM orders WHERE id = ?
		`, id)

	var items []*OrderItem
	err = d.db.SelectContext(ctx, &items, `
			SELECT * FROM order_items WHERE order_id = ?
			`, id)
	if err != nil {
		log.Println("Error getting order items", err)
		return nil, err
	}
	order.Items = items
	return &order, nil
}

func (d *MySQLStorer) ListOrders(ctx context.Context) ([]*Order, error) {
	var orders []*Order
	err := d.db.SelectContext(ctx, &orders, `
		SELECT * FROM orders
		`)
	if err != nil {
		log.Println("Error getting orders", err)
		return nil, err
	}
	return orders, nil
}

func (d *MySQLStorer) deleteOrderAndItems(ctx context.Context, tx *sqlx.Tx, orderId int64) error {
	_, err := tx.NamedExecContext(ctx, `
		DELETE FROM order_items WHERE order_id = ?
		`, orderId)
	if err != nil {
		log.Println("Error deleting order items", err)
		return err
	}
	_, err = tx.NamedExecContext(ctx, `
		DELETE FROM orders WHERE id = ?	
		`, orderId)
	if err != nil {
		log.Println("Error deleting order", err)
		return err
	}
	return nil
}

func (d *MySQLStorer) insertOrder(ctx context.Context, tx *sqlx.Tx, order *Order) (*Order, error) {
	res, err := tx.NamedExecContext(ctx, `
		INSERT INTO orders (payment_method, tax_price, shipping_price, total_price)
		VALUES (:payment_method, :tax_price, :shipping_price, :total_price)
		`, order)
	if err != nil {
		log.Println("Error inserting order", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id", err)
		return nil, err
	}

	order.ID = id

	return order, nil
}

func (d *MySQLStorer) insertOrderItem(ctx context.Context, tx *sqlx.Tx, orderItem *OrderItem) (*OrderItem, error) {
	res, err := tx.NamedExecContext(ctx, `
		INSERT INTO order_items (order_id, product_id, name, quantity, image, price)
		VALUES (:order_id, :product_id, :name, :quantity, :image, :price)
		`, orderItem)
	if err != nil {
		log.Println("Error inserting order item", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert id", err)
		return nil, err
	}

	orderItem.ID = id

	return orderItem, nil
}

func (d *MySQLStorer) execTx(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	tx, err := d.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Println("Error starting transaction", err)
		return err
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Println("Error rolling back transaction", rbErr)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction", err)
		return err
	}

	return nil
}
