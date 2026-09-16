package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abdelrahmanyasser2001/Go-CI-CD-Pipeline/src/models"
	"github.com/google/uuid"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func IDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, uuid.New().String())
}

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
