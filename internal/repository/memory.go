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

// NewMemoryRepository creates a new in-memory repository with some sample data
func NewMemoryRepository() *MemoryRepository {
	repo := &MemoryRepository{
		items:  make(map[int]models.Item),
		nextID: 1, // start IDs from 1
	}

	// --- Pre-populate with sample data for testing ---
	repo.items[repo.nextID] = models.Item{
		ID:          repo.nextID,
		Name:        "Sample Item 1",
		Description: "Sample Item 1 description",
		Quantity:    10,
	}
	repo.nextID++

	repo.items[repo.nextID] = models.Item{
		ID:          repo.nextID,
		Name:        "Sample Item 2",
		Description: "Sample Item 1 description",
		Quantity:    5,
	}
	repo.nextID++
	// -------------------------------------------------

	return repo
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
