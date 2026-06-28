package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/musishere/ecommerce-microservices/ecom-api/server"
	"github.com/musishere/ecommerce-microservices/ecom-api/storer"
)

type Handler struct {
	ctx    context.Context
	server *server.Server
}

func NewHandler(ctx context.Context, server *server.Server) *Handler {
	return &Handler{ctx: ctx, server: server}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// now := time.Now()
	created, err := h.server.CreateProduct(h.ctx, h.toStorerProduct(&req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // must come BEFORE Encode
	json.NewEncoder(w).Encode(h.toCreateProductResponse(created))
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.server.GetProduct(h.ctx, idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // must come BEFORE Encode
	json.NewEncoder(w).Encode(h.toCreateProductResponse(product))
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.server.ListProducts(h.ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	var resp []CreateProductResponse
	for _, product := range products {
		resp = append(resp, h.toCreateProductResponse(product))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // must come BEFORE Encode
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productToUpdate, err := h.server.GetProduct(h.ctx, idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//patch the product to the request
	patchProductToRequest(productToUpdate, &req)

	updated, err := h.server.UpdateProduct(h.ctx, productToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // must come BEFORE Encode
	json.NewEncoder(w).Encode(h.toCreateProductResponse(updated))

}

func (h *Handler) toStorerProduct(product *CreateProductRequest) *storer.Product {
	return &storer.Product{
		Name:         product.Name,
		Image:        product.Image,
		Category:     product.Category,
		Description:  product.Description,
		Rating:       product.Rating,
		NumReviews:   product.NumReviews,
		Price:        product.Price,
		CountInStock: product.CountInStock,
	}
}

func (h *Handler) toCreateProductResponse(product *storer.Product) CreateProductResponse {
	updatedAt := time.Time{}
	if product.UpdatedAt != nil {
		updatedAt = *product.UpdatedAt
	}
	return CreateProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Image:        product.Image,
		Category:     product.Category,
		Description:  product.Description,
		Rating:       product.Rating,
		NumReviews:   product.NumReviews,
		Price:        product.Price,
		CountInStock: product.CountInStock,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    updatedAt,
	}
}

func patchProductToRequest(product *storer.Product, req *CreateProductRequest) {
	if product.Name != "" {
		req.Name = product.Name
	}
	if product.Image != "" {
		req.Image = product.Image
	}
	if product.Category != "" {
		req.Category = product.Category
	}
	if product.Description != "" {
		req.Description = product.Description
	}
	if product.Rating != 0 {
		req.Rating = product.Rating
	}
	if product.NumReviews != 0 {
		req.NumReviews = product.NumReviews
	}
	if product.Price != 0 {
		req.Price = product.Price
	}
	if product.CountInStock != 0 {
		req.CountInStock = product.CountInStock
	}

	now := time.Now()
	product.UpdatedAt = &now
}
