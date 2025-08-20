package server

import (
	"net/http"
	"strconv"
	"strings"

	"inventory/inventory-api/internal/handlers"
	"inventory/inventory-api/internal/repository"
)

// RegisterRoutes sets up all API endpoints and attaches the repository
func RegisterRoutes(repo repository.Repository) {
	// Health check
	http.HandleFunc("/health", HealthHandler)

	// Items handler
	itemHandler := handlers.NewItemHandler(repo)

	// GET /items and POST /items
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

	// GET /items/{id} and DELETE /items/{id}
	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		// Extract ID from URL manually
		idStr := strings.TrimPrefix(r.URL.Path, "/items/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid item ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			itemHandler.GetItemByID(w, r, id)
		case http.MethodDelete:
			itemHandler.DeleteItem(w, r, id)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
