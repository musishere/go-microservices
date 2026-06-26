// Package classification of product api
//
// Documentation of product api
//
// Schemas: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
//   - application/json
//
// produces:
//   - application/json
// swagger:meta

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	protos "github.com/musishere/grpc/protos"
	"github.com/musishere/working-package/data"
)

type Products struct {
	l  *log.Logger
	cc protos.CurrencyClient
}

func NewProduct(l *log.Logger, cc protos.CurrencyClient) *Products {
	return &Products{l, cc}
}

func (p *Products) GetSingleProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, "Unable to convert id into string", http.StatusBadRequest)
		return
	}

	prod, _, err := data.FindProduct(idInt)
	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	err = prod.Tojson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	// get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies_USD,
		Destination: protos.Currencies_EUR,
	}

	resp, err := p.cc.GetRate(context.Background(), rr)
	p.l.Println("Response from currency service", resp)
	if err != nil {
		p.l.Println("Error getting exchange rate", err)
	}

	prod.Price = prod.Price * float64(resp.Rate)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Hanlde GET products")
	lp := data.GetProducts()
	err := lp.Tojson(rw)
	if err != nil {
		http.Error(rw, "internal servert error", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProducts(&prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, "Unable to convert id into string", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProducts(idInt, &prod)
	if err != nil {
		http.Error(rw, "Unable to update product", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct {
}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			fmt.Println("validation error", err)
			http.Error(rw, "Validation Failed", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
