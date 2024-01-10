package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ExampleData represents a data structure to be returned by the API
type ExampleData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// getExampleData handles GET requests and returns JSON data
func getExampleData(w http.ResponseWriter, r *http.Request) {
	// Set content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Creating a response data
	data := ExampleData{
		ID:   "1",
		Name: "Mux",
	}

	// Encode data to JSON and send as a response
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define endpoints
	router.HandleFunc("/api/data", getExampleData).Methods("GET")

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
