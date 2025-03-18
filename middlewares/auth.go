package middlewares

import (
	"app/repository"
	"log/slog"
	"net/http"
)

type AuthMiddleware struct {
	sessionRepo *repository.SessionRepo
	logger      *slog.Logger
}

func NewAuthMiddleware(
	sessionRepo *repository.SessionRepo,
	logger *slog.Logger,
) *AuthMiddleware {
	return &AuthMiddleware{
		sessionRepo: sessionRepo,
		logger:      logger,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session, err := m.sessionRepo.GetByID(sessionCookie.Value)
		if err != nil {
			m.logger.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		if session == nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
