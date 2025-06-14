package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"server-database/db"
)

// ImagesHandler handles both GET and POST requests for images
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetImages(w, r)
	case http.MethodPost:
		handlePostImages(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetImages(w http.ResponseWriter, r *http.Request) {
	// Query database for images
	images, err := db.GetAllImages(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check for "indent" query parameter
	indentParam := r.URL.Query().Get("indent")
	indent := 0
	if indentParam != "" {
		if i, err := strconv.Atoi(indentParam); err == nil && i >= 0 && i <= 10 {
			indent = i
		}
	}

	// Serialize data to JSON
	var imagesData []byte
	var jsonErr error // Declare the error variable
	if indent > 0 {
		imagesData, jsonErr = json.MarshalIndent(images, "", strings.Repeat(" ", indent))
	} else {
		imagesData, jsonErr = json.Marshal(images)
	}
	if jsonErr != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(imagesData)
}

func handlePostImages(w http.ResponseWriter, r *http.Request) {
	// FOR POST logic
}
