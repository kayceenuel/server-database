package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// Define the static JSON data
	data := map[string]interface{}{
		"message": "Hello, World!",
		"status":  "success",
		"data": map[string]string{
			"author": "John Doe",
			"year":   "2025",
		},
	}

	// HTTP handler function
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
