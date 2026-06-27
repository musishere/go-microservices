package main

import (
	"fmt"
	"log"

	db "github.com/musishere/ecommerce-microservices/db/migrations"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error opening database")
		return
	}

	defer db.Close()
	fmt.Println("Succesfully connected to database")

}
