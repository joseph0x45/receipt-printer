package main

import (
	"app/handlers"
	"app/helpers"
	"app/repository"
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

//go:embed views/*.html
var templatesFS embed.FS

func newLogger() *slog.Logger {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
		}),
	)
	return logger
}

func main() {
	logger := newLogger()
	godotenv.Load()
	port := helpers.GetEnvVar("PORT", "8080")

	//repositories
	productRepo := repository.NewProductRepo(nil)

	//handlers
	productHandler := handlers.NewProductHandler(
		productRepo,
		&templatesFS,
		logger,
	)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	productHandler.RegisterRoutes(r)

	server := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Println("Starting server on port", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
