package main

import (
	"encoding/json"
	"net/http"
)

// Image represents the structure of the image data.
type Image struct {
	Title   string `json:"title"`
	AltText string `json:"alt_text"`
	URL     string `json:"url"`
}

func main() {
	http.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		// initialising image data
		images := []Image{
			{
				Title:   "Sunset",
				AltText: "Clouds at sunset",
				URL:     "https://images.unsplash.com/photo-1506815444479-bfdb1e96c566?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80",
			},
			{
				Title:   "Mountain",
				AltText: "A mountain at sunset",
				URL:     "https://images.unsplash.com/photo-1540979388789-6cee28a1cdc9?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1000&q=80",
			},
		}
		//response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Encode and send JSON response
		json.NewEncoder(w).Encode(images)
	})
	// start the http server
	http.ListenAndServe(":8080", nil)
}
