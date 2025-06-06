package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Image represents the structure of the image data.
type Image struct {
	Title   string `json:"title"`
	AltText string `json:"alt_text"`
	URL     string `json:"url"`
}

func handleImages(w http.ResponseWriter, r *http.Request) {
	// image data in json format
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

	// check for "indent" query parameter
	indentParam := r.URL.Query().Get("indent")
	indent := 0
	if indentParam != "" {
		if i, err := strconv.Atoi(indentParam); err == nil && i >= 0 && i <= 10 {
			indent = i
		}
	}

	// serialize data to Json
	var imagesData []byte
	var err error
	if indent > 0 {
		imagesData, err = json.MarshalIndent(images, "", strings.Repeat(" ", indent))
	} else {
		imagesData, err = json.Marshal(images)
	}
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	// set response headers and write JSON
	w.Header().Set("Content-Type", "text/json")
	w.Write(imagesData)
}

func main() {
	http.HandleFunc("/images.json", handleImages)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
