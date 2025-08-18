package server

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler handles GET /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := HealthResponse{Status: "ok"}
	json.NewEncoder(w).Encode(resp)
}
