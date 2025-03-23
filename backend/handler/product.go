package handler

import (
	"backend/repository"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	productRepo *repository.ProductRepo
	logger      *slog.Logger
}

func NewProductHandler(
	productRepo *repository.ProductRepo,
	logger *slog.Logger,
) *ProductHandler {
	return &ProductHandler{
		productRepo: productRepo,
		logger:      logger,
	}
}

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) getAllProducts(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) RegisterRoutes(r chi.Router) {
	r.Post("/products", h.createProduct)
	r.Get("/products", h.getAllProducts)
}
