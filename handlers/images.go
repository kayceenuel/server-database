package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"server-database/db"
	"server-database/models"
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
	images, err := db.GetAllImages(r.Context())
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
	var img models.Image
	if err := json.NewDecoder(r.Body).Decode(&img); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// validate the image
	if err := img.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check for duplicate URL
	exists, err := db.CheckDuplicateURL(r.Context(), img.URL)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if exists {
		http.Error(w, "Image with this URL already exists", http.StatusConflict)
		return
	}

	// Insert into database
	newImage, err := db.InsertImage(r.Context(), img)
	if err != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newImage)
}
