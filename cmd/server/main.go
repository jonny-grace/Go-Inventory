package main

import (
	"fmt"
	"log"
	"net/http"

	"inventory/inventory-api/internal/server"
)

func main() {
	port := 8080

	// Register routes from internal/server
	server.RegisterRoutes()

	log.Printf("🚀 Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
