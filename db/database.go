package db

import (
	"context"
	"fmt"
	"os"
	"server-database/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

// Connect establishes a connection pool to PostgreSQL
func Connect() error {
	// Load the connection string from environment variable
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		return fmt.Errorf("DATABASE_URL environment variable not set")
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Initialize the connection pool
	var err error
	Pool, err = pgxpool.New(ctx, connString)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %v", err)
	}

	// Verify the connection
	if err := Pool.Ping(ctx); err != nil {
		Pool.Close() // Clean up on ping failure
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to database")
	return nil
}

// Close terminates the connection pool
func Close() {
	if Pool != nil {
		Pool.Close()
		fmt.Println("Database connection pool closed")
	}
}

// GetAllImages retrieves all images from the database
func GetAllImages(ctx context.Context) ([]models.Image, error) {
	rows, err := Pool.Query(ctx, "SELECT id, title, alt_text, url FROM images")
	if err != nil {
		return nil, fmt.Errorf("failed to query images: %v", err)
	}
	defer rows.Close()

	var images []models.Image
	for rows.Next() {
		var img models.Image
		if err := rows.Scan(&img.ID, &img.Title, &img.AltText, &img.URL); err != nil {
			return nil, fmt.Errorf("failed to scan image: %v", err)
		}
		images = append(images, img)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return images, nil
}

// InsertImage adds a new image and returns the created image with ID
func InsertImage(ctx context.Context, img models.Image) (*models.Image, error) {
	var newImage models.Image
	err := Pool.QueryRow(ctx,
		"INSERT INTO images (title, alt_text, url) VALUES ($1, $2, $3) RETURNING id, title, alt_text, url",
		img.Title, img.AltText, img.URL).Scan(&newImage.ID, &newImage.Title, &newImage.AltText, &newImage.URL)

	if err != nil {
		return nil, fmt.Errorf("failed to insert image: %v", err)
	}
	return &newImage, nil
}
