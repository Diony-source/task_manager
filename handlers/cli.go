package handlers

import (
	"fmt"
	"task_manager/services"
)

// StartCLI initializes the command-line interface for task management
func StartCLI() {
	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. Update Task")
		fmt.Println("3. Complete Task")
		fmt.Println("4. List Tasks")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task title: ")
			var title string
			fmt.Scanln(&title)

			fmt.Print("Enter task description: ")
			var description string
			fmt.Scanln(&description)

			err := services.AddTask(title, description)
			if err != nil {
				fmt.Println("Error adding task:", err)
			} else {
				fmt.Println("Task added successfully!")
			}

		case 2:
			fmt.Print("Enter task ID to update: ")
			var id int
			fmt.Scanln(&id)

			fmt.Print("Enter new task title: ")
			var title string
			fmt.Scanln(&title)

			fmt.Print("Enter new task description: ")
			var description string
			fmt.Scanln(&description)

			err := services.UpdateTask(id, title, description)
			if err != nil {
				fmt.Println("Error updating task:", err)
			} else {
				fmt.Println("Task updated successfully!")
			}

		case 3:
			fmt.Print("Enter task ID to mark as completed: ")
			var id int
			fmt.Scanln(&id)

			err := services.CompleteTask(id)
			if err != nil {
				fmt.Println("Error completing task:", err)
			} else {
				fmt.Println("Task marked as completed!")
			}

		case 4:
			tasks, err := services.ListTasks()
			if err != nil {
				fmt.Println("Error listing tasks:", err)
			} else {
				for _, task := range tasks {
					status := "Incomplete"
					if task.IsCompleted {
						status = "Completed"
					}
					fmt.Printf("ID: %d | Title: %s | Description: %s | Status: %s\n",
						task.ID, task.Title, task.Description, status)
				}
			}

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
