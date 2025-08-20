package handlers

import (
	"encoding/json"
	"inventory/inventory-api/internal/models"
	"inventory/inventory-api/internal/repository"
	"net/http"
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

// GetItemByID handles GET /items/{id}
func (h *ItemHandler) GetItemByID(w http.ResponseWriter, r *http.Request, id int) {
	item, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// CreateItem handles POST /items
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(&item); err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// DeleteItem handles DELETE /items/{id}
func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request, id int) {
	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdateItem handles PUT /items/{id}
func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request, id int) {
	var updatedItem models.Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get existing item first
	existingItem, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	// Update fields
	existingItem.Name = updatedItem.Name
	existingItem.Description = updatedItem.Description
	existingItem.Quantity = updatedItem.Quantity
	existingItem.Price = updatedItem.Price

	// Save updated item
	if err := h.Repo.Update(existingItem); err != nil {
		http.Error(w, "Failed to update item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingItem)
}
