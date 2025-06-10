package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool holds the db connection pool
var Pool *pgxpool.Pool

func Connect() error {
	//load the connection string from the .env var
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		return fmt.Errorf("DATABASE_URL environment Variable not set")
	}

	// Initialize the connection pool
	var err error
	Pool, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %v", err)
	}
	//verify the connection with context
	if err := Pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("failed to ping databse: %v", err)
	}
	fmt.Println("Successfully created connection Pool")
}

// Close terminates the connection Pool
func Close() {
	if Pool != nil {
		Pool.Close()
		fmt.Println("Connection pool Closed")
	}
}
