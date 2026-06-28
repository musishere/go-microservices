package main

import (
	"fmt"
	"log"

	db "github.com/musishere/ecommerce-microservices/db/migrations"
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
	_ := server.NewServer(storer)

	// handler := handler.NewHandler(srvr)
}
