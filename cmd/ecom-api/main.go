package main

import (
	"flag"
	"fmt"
	"log"

	db "github.com/musishere/ecommerce-microservices/db/migrations"
	"google.golang.org/grpc"
)

func main() {
	var (
		grpcAddr = flag.String("grpc-addr", "localhost:50051", "The address of the grpc server")
	)
	flag.Parse()

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database")
		return
	}

	defer db.Close()
	fmt.Println("Succesfully connected to database")

	grpc.NewClient(*grpcAddr)

}
