package repository

import "inventory/inventory-api/internal/models"

// Repository defines the behavior for storing and retrieving items
type Repository interface {
	GetAll() ([]models.Item, error)
	GetByID(id int) (*models.Item, error)
	Create(item *models.Item) error
	Delete(id int) error
}
