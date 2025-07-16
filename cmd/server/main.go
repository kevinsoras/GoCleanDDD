package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kevinsoras/GoCleanDDD/internal/db/postgres"
	silicon "github.com/kevinsoras/GoCleanDDD/internal/silicon"
	middleware "github.com/kevinsoras/GoCleanDDD/pkg/middleware"
)

func main() {
	router := http.NewServeMux()
	// Cargar .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create middleware stack
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	// Create database connection
	db, err := postgres.NewPostgresPool(context.Background())
	if err != nil {
		log.Fatal("Error creating database connection pool", err)
	}
	defer db.Close()
	// Load routes
	silicon := silicon.NewSiliconModule(db)
	silicon.RegisterRoutes(router)

	log.Println("Server listening on :8080")
	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	server.ListenAndServe()
}
