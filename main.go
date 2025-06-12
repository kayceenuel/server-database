package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"server-database/db"

	"github.com/joho/godotenv"
)

// Image represents the structure of the image data.
type Image struct {
	Title   string `json:"title"`
	AltText string `json:"alt_text"`
	URL     string `json:"url"`
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	// Query database for images
	rows, err := db.Pool.Query(context.Background(), "SELECT title, url, alt_text FROM images")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Read the results into a slice
	var images []Image
	for rows.Next() {
		var img Image
		err := rows.Scan(&img.Title, &img.URL, &img.AltText)
		if err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		images = append(images, img)
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		http.Error(w, "Error reading database rows", http.StatusInternalServerError)
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
	if indent > 0 {
		imagesData, err = json.MarshalIndent(images, "", strings.Repeat(" ", indent))
	} else {
		imagesData, err = json.Marshal(images)
	}
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(imagesData)
}

func main() {
	// load environment variable from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No, .env file found, using system enviroment variables")
	}
	// Initialize the connection pool
	if err := db.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close() // Close the pool connection when the app exits

	// HTTP Handler
	http.HandleFunc("/images.json", imagesHandler)

	// Start the server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
