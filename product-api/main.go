package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	protos "github.com/musishere/grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gorilla/mux"
	"github.com/musishere/working-package/handlers"
)

func main() {
	l := log.New(os.Stdout, "Product-Api", log.LstdFlags)

	// create a new client for grpc
	conn, err := grpc.NewClient(":8010", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		l.Fatal("Error while making grpc client server")
	}
	defer conn.Close()
	cc := protos.NewCurrencyClient(conn)

	ph := handlers.NewProduct(l, cc)

	sm := mux.NewRouter()

	// GET routes
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)
	getRouter.HandleFunc("/{id:[0-9]+}", ph.GetSingleProduct)

	// PUT routes
	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.Use(ph.MiddlewareProductValidation)
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)

	// POST handler
	postRouter := sm.Methods("POST").Subrouter()
	postRouter.Use(ph.MiddlewareProductValidation)
	postRouter.HandleFunc("/", ph.AddProduct)

	s := &http.Server{
		Addr:         ":2080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		fmt.Println("server running on port 2080")
		err := s.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("Received signal:", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(tc)
}
