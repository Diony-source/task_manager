package main

import (
	"task_manager/database"
	"task_manager/handlers"
)

// Main entry point for the Task Manager application
func main() {
	// Initialize the database
	database.InitDB()

	// Start the CLI
	handlers.StartCLI()
}
