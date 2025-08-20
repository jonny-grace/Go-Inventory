package main

import (
	"fmt"
	"log"
	"net/http"

	"inventory/inventory-api/internal/repository"
	"inventory/inventory-api/internal/server"
)

func main() {
	useDB := false // set true to use Postgres

	var repo repository.Repository
	var err error

	if useDB {
		connString := "postgres://jgrace:1234567890@localhost:5432/inventorytest?sslmode=disable"
		repo, err = repository.NewPostgresRepository(connString)
		if err != nil {
			log.Fatalf("Failed to connect to Postgres: %v", err)
		}
	} else {
		repo = repository.NewMemoryRepository()
	}

	// Pass the repository to RegisterRoutes
	server.RegisterRoutes(repo)

	port := 8080
	log.Printf("🚀 Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
