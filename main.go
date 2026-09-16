package main

import (
	"fmt"
	"net/http"

	"github.com/abdelrahmanyasser2001/Go-CI-CD-Pipeline/src/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/health", handlers.HealthHandler)
	r.HandleFunc("/id", handlers.IDHandler)

	fmt.Println("listening on :3000")
	http.ListenAndServe(":3000", r)
}
