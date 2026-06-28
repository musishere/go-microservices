package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	db "github.com/musishere/ecommerce-microservices/db/migrations"
	"github.com/musishere/ecommerce-microservices/ecom-grpc/pb"
	grpcserver "github.com/musishere/ecommerce-microservices/ecom-grpc/server"
	"github.com/musishere/ecommerce-microservices/ecom-grpc/storer"
	"google.golang.org/grpc"
)

func main() {
	var (
		port = flag.Int("port", 50051, "The server port")
		addr = fmt.Sprintf(":%d", *port)
	)
	flag.Parse()

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database")
		return
	}

	defer db.Close()
	fmt.Println("Succesfully connected to database")

	storer := storer.NewMySQLStorer(db.GetDB())
	srvr := grpcserver.NewServer(storer)

	grpcServer := grpc.NewServer()
	pb.RegisterEcommServer(grpcServer, srvr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Error listening: %v", err)
		return
	}

	fmt.Println("Server listening on port", *port)
	grpcServer.Serve(listener)
	if err != nil {
		log.Printf("Error serving: %v", err)
		return
	}

}
