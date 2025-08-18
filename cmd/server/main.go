package main

import (
	"fmt"
	"log"
	"net/http"

	"inventory/inventory-api/internal/repository"
	"inventory/inventory-api/internal/server"
)

func main() {
	// Initialize in-memory repository
	repo := repository.NewMemoryRepository()

	port := 8080

	// Register routes (we pass repo if needed later for CRUD)
	server.RegisterRoutes(repo)

	log.Printf("🚀 Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil)))
}