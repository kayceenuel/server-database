package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool holds the database connection pool
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
