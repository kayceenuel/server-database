package main

import (
	"log"
	"net/http"

	"server-database/db"
	"server-database/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database connection
	if err := db.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Set up routes
	http.HandleFunc("/images.json", handlers.ImagesHandler)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
