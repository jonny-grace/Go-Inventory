package server

import (
	"net/http"

	"inventory/inventory-api/internal/handlers"
	"inventory/inventory-api/internal/repository"
)

// RegisterRoutes sets up all API endpoints and attaches the repository
func RegisterRoutes(repo repository.Repository) {
	// Health check
	http.HandleFunc("/health", HealthHandler)

	// Items handler
	itemHandler := handlers.NewItemHandler(repo)

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			itemHandler.GetAllItems(w, r)
		case http.MethodPost:
			itemHandler.CreateItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/items/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			itemHandler.GetItemByID(w, r)
		case http.MethodDelete:
			itemHandler.DeleteItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
