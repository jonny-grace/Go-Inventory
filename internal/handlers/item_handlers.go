package handlers

import (
	"encoding/json"
	"net/http"

	"inventory/inventory-api/internal/models"
	"inventory/inventory-api/internal/repository"
)

type ItemHandler struct {
	Repo repository.Repository
}

// NewItemHandler creates a new handler with the given repository
func NewItemHandler(repo repository.Repository) *ItemHandler {
	return &ItemHandler{Repo: repo}
}

// GetAllItems handles GET /items
func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// CreateItem handles POST /items
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item *models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(item); err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
