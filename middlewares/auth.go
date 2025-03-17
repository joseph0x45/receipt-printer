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
		sessionID := r.Header.Get("Authorization")
		if sessionID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		session, err := m.sessionRepo.GetByID(sessionID)
		if err != nil {
			m.logger.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
