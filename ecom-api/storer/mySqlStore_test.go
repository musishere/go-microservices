package storer

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func WithTestDB(t *testing.T, fn func(*sqlx.DB, sqlmock.Sqlmock)) {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("Error creating sqlmock: %v", err)
	}
	defer mockDb.Close()
	db := sqlx.NewDb(mockDb, "mysql")
	fn(db, mock)
}

func TestCreateProduct(t *testing.T) {
	WithTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		storer := NewMySQLStorer(db)

		p := &Product{Name: "Test Product", Image: "test.jpg", Category: "Test Category", Description: "Test Description", Rating: 5, NumReviews: 10, Price: 100, CountInStock: 10}

		mock.ExpectExec("INSERT INTO products (name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES (?, ?, ?, ?, ?, ?, ?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))

		product, err := storer.CreateProduct(context.Background(), p)
		require.NoError(t, err)
		require.Equal(t, int64(1), product.ID)
		err = mock.ExpectationsWereMet()
		require.NoError(t, err)
	})
}

func TestGetProductByID(t *testing.T) {
	WithTestDB(t, func(db *sqlx.DB, mock sqlmock.Sqlmock) {
		storer := NewMySQLStorer(db)

		rows := sqlmock.NewRows([]string{"id", "name", "image", "category", "description", "rating", "num_reviews", "price", "count_in_stock"}).AddRow(1, "Test Product", "test.jpg", "Test Category", "Test Description", 5, 10, 100, 10)
		mock.ExpectQuery("SELECT * FROM products WHERE id = ?").WithArgs(1).WillReturnRows(rows)

		product, err := storer.GetProductByID(context.Background(), 1)
		require.NoError(t, err)
		require.Equal(t, int64(1), product.ID)

	})
}
