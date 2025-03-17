package main

import (
	"app/handlers"
	"app/helpers"
	"app/middlewares"
	"app/repository"
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
  _ "github.com/lib/pq"
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

func getDBPool() *sqlx.DB {
	dbURL := os.Getenv("DB_URL")
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
  log.Println("Connected to database")
	return db
}

func main() {
	logger := newLogger()
	godotenv.Load()
	port := helpers.GetEnvVar("PORT", "8080")

	db := getDBPool()
	if db == nil {
		panic("Failed to connect to database")
	}

	//repositories
	productRepo := repository.NewProductRepo(db)
	sessionRepo := repository.NewSessionRepo(db)

	//middlewares
	authMiddleware := middlewares.NewAuthMiddleware(
		sessionRepo,
		logger,
	)

	//handlers
	productHandler := handlers.NewProductHandler(
		productRepo,
		&templatesFS,
		logger,
		authMiddleware,
	)

	authHandler := handlers.NewAuthHandler(
		sessionRepo,
		logger,
		&templatesFS,
	)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	})

	productHandler.RegisterRoutes(r)
	authHandler.RegisterRoutes(r)

	server := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Println("Starting server on port", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
