package server

import (
	"net/http"

	"inventory/inventory-api/internal/repository"
)

// RegisterRoutes sets up all API endpoints and attaches the repository
func RegisterRoutes(repo repository.Repository) {
	// Health check
	http.HandleFunc("/health", HealthHandler)

	// Later we’ll add item routes (GET, POST, etc.)
	// Example:
	// http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
	// 	ItemsHandler(w, r, repo)
	// })
}