package main

import (
	"log"
	"task_manager/database"
	"task_manager/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database
	database.InitDB()

	// Start the CLI
	handlers.StartCLI()
}
