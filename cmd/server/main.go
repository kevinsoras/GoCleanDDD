package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
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

	silicon.LoadRoutesSilicon(router)

	log.Println("Server listening on :8080")
	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}
	server.ListenAndServe()
}
