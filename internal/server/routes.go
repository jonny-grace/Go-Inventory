package server

import (
	"net/http"
)

// RegisterRoutes sets up all HTTP routes for the app
func RegisterRoutes() {
	http.HandleFunc("/health", HealthHandler)
}
