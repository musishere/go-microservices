package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/musishere/working-package/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
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

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
		return
	}

	data.AddProducts(prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, "Unable to convert id into string", http.StatusBadRequest)
		return
	}
	prod := &data.Product{}
	err = prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
		return
	}

	err = data.UpdateProducts(idInt, prod)
	if err != nil {
		http.Error(rw, "Unable to update product", http.StatusInternalServerError)
		return
	}
}
