package handlers

import (
	"app/middlewares"
	"app/repository"
	"embed"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
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

func (h *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(h.viewsFS, "views/layout.html", "views/create-product.html")
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

func (h *ProductHandler) RegisterRoutes(r chi.Router) {
	r.With(h.authMiddleware.Authenticate).
		Get("/admin/create-product", h.createProduct)
	r.Get("/products", h.renderProductsPage)
}
