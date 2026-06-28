package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(h *Handler) *chi.Mux {

	r := chi.NewRouter()

	r.Route("/products", func(r chi.Router) {
		r.Post("/", h.CreateProduct)
		r.Get("/", h.ListProducts)
		r.Get("/{id}", h.GetProduct)
		r.Patch("/{id}", h.UpdateProduct)
		// r.Delete("/{id}", h.DeleteProduct)
	})

	return r
}

func StartServer(addr string, r *chi.Mux) error {
	return http.ListenAndServe(addr, r)
}
