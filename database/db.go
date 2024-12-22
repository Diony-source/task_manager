package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Global variable to hold the database connection pool
var DB *pgxpool.Pool

// InitDB initializes the PostgreSQL connection and creates the tasks table if it does not exist
func InitDB() {
	// Retrieve PostgreSQL connection string from environment variables
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set in environment variables")
	}

	// Connect to the PostgreSQL database
	var err error
	DB, err = pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Create tasks table if it does not exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		is_completed BOOLEAN NOT NULL DEFAULT FALSE
	);`
	_, err = DB.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create tasks table: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL database successfully!")
}
