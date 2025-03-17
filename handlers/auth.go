package handlers

import (
	"app/models"
	"app/repository"
	"app/shared"
	"embed"
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	sessionRepo *repository.SessionRepo
	logger      *slog.Logger
	viewsFS     *embed.FS
}

func NewAuthHandler(
	sessionRepo *repository.SessionRepo,
	logger *slog.Logger,
	viewsFS *embed.FS,
) *AuthHandler {
	return &AuthHandler{
		sessionRepo: sessionRepo,
		logger:      logger,
		viewsFS:     viewsFS,
	}
}

func (h *AuthHandler) renderLoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(h.viewsFS, "views/layout.html", "views/login.html")
	if err != nil {
		h.logger.Error("Error loading template:" + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pageData := struct {
		Title string
	}{
		Title: "Login",
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		h.logger.Error("Error while executing template: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func (h *AuthHandler) processLogin(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if payload.Password == "" || len(payload.Password) >= 72 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(shared.PASSWORD), []byte(payload.Password)) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	session := &models.Session{
		ID:     ulid.Make().String(),
		Active: true,
	}
	err = h.sessionRepo.Insert(session)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name:     "session",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) RegisterRoutes(r chi.Router) {
	r.Get("/login", h.renderLoginPage)
	r.Post("/api/login", h.processLogin)
}
