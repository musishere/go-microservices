package main

import (
	"context"
	"fmt"
	"log"

	db "github.com/musishere/ecommerce-microservices/db/migrations"
	"github.com/musishere/ecommerce-microservices/ecom-api/handler"
	"github.com/musishere/ecommerce-microservices/ecom-api/server"
	"github.com/musishere/ecommerce-microservices/ecom-api/storer"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database")
		return
	}

	defer db.Close()
	fmt.Println("Succesfully connected to database")

	storer := storer.NewMySQLStorer(db.GetDB())
	srvr := server.NewServer(storer)
	h := handler.NewHandler(context.Background(), srvr)
	r := handler.RegisterRoutes(h)
	err = handler.StartServer(":8080", r)
	if err != nil {
		log.Fatal("Error starting server")
		return
	}
}
