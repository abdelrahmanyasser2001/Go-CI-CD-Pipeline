package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YOUR_USERNAME/go-mini/src/models"
	"github.com/google/uuid"
)

// HomeHandler handles GET /
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// IDHandler handles GET /id, returning a freshly generated UUID.
func IDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, uuid.New().String())
}

// HealthHandler handles GET /health, used by the pipeline/staging check
// to confirm the deployed build is actually up and responding.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := models.Response{
		Message: "I'm alive!",
		Status:  http.StatusOK,
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
