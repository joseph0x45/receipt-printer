package handlers

import (
	"app/repository"
	"embed"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	productsRepo *repository.ProductRepo
	viewsFS      *embed.FS
	logger       *slog.Logger
}

func NewProductHandler(
	productsRepo *repository.ProductRepo,
	viewsFS *embed.FS,
	logger *slog.Logger,
) *ProductHandler {
	return &ProductHandler{
		productsRepo: productsRepo,
		viewsFS:      viewsFS,
		logger:       logger,
	}
}

func (h *ProductHandler) renderProductsPage(w http.ResponseWriter, r *http.Request) {
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

func (h *ProductHandler) RegisterRoutes(r chi.Router) {
	r.Get("/products", h.renderProductsPage)
}
