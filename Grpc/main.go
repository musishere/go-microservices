package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	currency "github.com/musishere/grpc/protos"
	"github.com/musishere/grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Register a Logger
	log := hclog.Default()

	// Invoke a Service
	cs := server.NewCurrency(log)

	// Initialize a grpc server
	gs := grpc.NewServer()

	// Register grpc and service server
	currency.RegisterCurrencyServer(gs, cs)

	// Expose the grpc server
	reflection.Register(gs)

	// Register a listner
	l, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Error("Error on making listner", err)
		os.Exit(1)
	}

	// Start the server
	gs.Serve(l)
}
