package handlers

import (
	"app/middlewares"
	"app/models"
	"app/repository"
	"embed"
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
)

type ProductHandler struct {
	productsRepo   *repository.ProductRepo
	viewsFS        *embed.FS
	logger         *slog.Logger
	authMiddleware *middlewares.AuthMiddleware
}

func NewProductHandler(
	productsRepo *repository.ProductRepo,
	viewsFS *embed.FS,
	logger *slog.Logger,
	authMiddleware *middlewares.AuthMiddleware,
) *ProductHandler {
	return &ProductHandler{
		productsRepo:   productsRepo,
		viewsFS:        viewsFS,
		logger:         logger,
		authMiddleware: authMiddleware,
	}
}

func (h *ProductHandler) renderProductsPage(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(h.viewsFS, "views/layout.html", "views/products.html")
	if err != nil {
		h.logger.Error("Error loading template:" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pageData := struct {
		Title string
	}{
		Title: "Products",
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		h.logger.Error("Error while executing template: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func (h *ProductHandler) renderAddProductPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(h.viewsFS, "views/layout.html", "views/add-product.html")
	if err != nil {
		h.logger.Error("Error loading template:" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pageData := struct {
		Title string
	}{
		Title: "Add Product",
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		h.logger.Error("Error while executing template: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func (h *ProductHandler) processProductCreation(w http.ResponseWriter, r *http.Request) {
	payload := &struct {
		Name      string `json:"name"`
		Media     string `json:"media"`
		Price     int    `json:"price"`
		BulkPrice int    `json:"bulk_price"`
		InStock   int    `json:"in_stock"`
	}{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newProduct := &models.Product{
		ID:        ulid.Make().String(),
		Name:      payload.Name,
		Media:     payload.Media,
		Price:     payload.Price,
		BulkPrice: payload.BulkPrice,
		InStock:   payload.InStock,
	}
	err = h.productsRepo.Insert(newProduct)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) RegisterRoutes(r chi.Router) {
	r.With(h.authMiddleware.Authenticate).Get("/add-product", h.renderAddProductPage)
	r.With(h.authMiddleware.Authenticate).Post("/api/products", h.processProductCreation)
	r.Get("/products", h.renderProductsPage)
}
