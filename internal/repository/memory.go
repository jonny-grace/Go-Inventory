package repository

import (
	"errors"
	"inventory/inventory-api/internal/models"
	"sync"
)

type MemoryRepository struct {
	mu     sync.Mutex
	items  map[int]models.Item
	nextID int
}

// NewMemoryRepository creates a new in-memory repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		items: make(map[int]models.Item),
	}
}

func (r *MemoryRepository) Create(item *models.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	item.ID = r.nextID
	r.items[item.ID] = *item
	r.nextID++
	return nil
}

func (r *MemoryRepository) GetAll() ([]models.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	items := make([]models.Item, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}
	return items, nil
}

func (r *MemoryRepository) GetByID(id int) (*models.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.items[id]
	if !exists {
		return nil, errors.New("item not found")
	}
	return &item, nil
}

func (r *MemoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.items[id]; !exists {
		return errors.New("item not found")
	}
	delete(r.items, id)
	return nil
}
