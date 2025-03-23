package main

import (
	"backend/handler"
	"backend/repository"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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
		log.Println("Failed to connect to database:", err.Error())
		return nil
	}
	log.Println("Connected to database")
	return db
}

func main() {
	godotenv.Load()
	logger := newLogger()
	dbPool := getDBPool()
	if dbPool == nil {
		return
	}

	//repositories
	productRepo := repository.NewProductRepo(dbPool)

	//handlers
	productHandler := handler.NewProductHandler(
		productRepo,
		logger,
	)

	r := chi.NewRouter()

	productHandler.RegisterRoutes(r)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: time.Minute,
		ReadTimeout:  time.Minute,
		IdleTimeout:  time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}
