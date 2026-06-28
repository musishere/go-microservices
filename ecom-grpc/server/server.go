package server

import (
	"context"

	"github.com/musishere/ecommerce-microservices/ecom-grpc/pb"
	"github.com/musishere/ecommerce-microservices/ecom-grpc/storer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	storer *storer.MySQLStorer
	pb.UnimplementedEcommServer
}

func NewServer(storer *storer.MySQLStorer) *Server {
	return &Server{storer: storer}
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.ProductReq) (*pb.ProductRes, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateProduct not implemented")
}
