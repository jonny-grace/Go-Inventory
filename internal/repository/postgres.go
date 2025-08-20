package repository

import (
	"context"
	"errors"
	"inventory/inventory-api/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository connects to Postgres
func NewPostgresRepository(connString string) (*PostgresRepository, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: pool}, nil
}

// GetAll returns all items
func (r *PostgresRepository) GetAll() ([]models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, "SELECT id, name, description, quantity, price FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []models.Item{}
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// GetByID returns a single item by ID
func (r *PostgresRepository) GetByID(id int) (*models.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var item models.Item
	err := r.db.QueryRow(ctx, "SELECT id, name, description, quantity, price FROM items WHERE id=$1", id).
		Scan(&item.ID, &item.Name, &item.Description, &item.Quantity, &item.Price)
	if err != nil {
		return nil, errors.New("item not found")
	}

	return &item, nil
}

// Create inserts a new item
func (r *PostgresRepository) Create(item *models.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.db.QueryRow(
		ctx,
		"INSERT INTO items (name, description, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id",
		item.Name, item.Description, item.Quantity, item.Price,
	).Scan(&item.ID)
}

// Delete removes an item by ID
func (r *PostgresRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ct, err := r.db.Exec(ctx, "DELETE FROM items WHERE id=$1", id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("item not found")
	}

	return nil
}

// Update updates an item by ID
func (r *PostgresRepository) Update(item *models.Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ct, err := r.db.Exec(ctx, "UPDATE items SET name=$1, description=$2, quantity=$3, price=$4 WHERE id=$5",
		item.Name, item.Description, item.Quantity, item.Price, item.ID)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return errors.New("item not found")
	}
	return nil
}
