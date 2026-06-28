package server

import (
	"context"

	"github.com/musishere/ecommerce-microservices/ecom-api/storer"
)

type Server struct {
	storer *storer.MySQLStorer
}

func NewServer(storer *storer.MySQLStorer) *Server {
	return &Server{storer: storer}
}

func (s *Server) CreateProduct(ctx context.Context, product *storer.Product) (*storer.Product, error) {
	return s.storer.CreateProduct(ctx, product)
}

func (s *Server) GetProduct(ctx context.Context, id int64) (*storer.Product, error) {
	return s.storer.GetProductByID(ctx, id)
}

func (s *Server) ListProducts(ctx context.Context) ([]*storer.Product, error) {
	return s.storer.ListProducts(ctx)
}

func (s *Server) UpdateProduct(ctx context.Context, product *storer.Product) (*storer.Product, error) {
	return s.storer.UpdateProduct(ctx, product)
}

func (s *Server) DeleteProduct(ctx context.Context, id int64) error {
	return s.storer.DeleteProduct(ctx, id)
}

func (s *Server) CreateOrder(ctx context.Context, order *storer.Order) (*storer.Order, error) {
	return s.storer.CreateOrder(ctx, order)
}

func (s *Server) GetOrderByID(ctx context.Context, id int64) (*storer.Order, error) {
	return s.storer.GetOrderByID(ctx, id)
}
