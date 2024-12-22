package services

import (
	"context"
	"task_manager/database"
	"task_manager/entities"
)

// AddTask inserts a new task into the database
func AddTask(title, description string) error {
	query := "INSERT INTO tasks (title, description) VALUES ($1, $2)"
	_, err := database.DB.Exec(context.Background(), query, title, description)
	return err
}

// UpdateTask modifies the title and description of an existing task
func UpdateTask(id int, title, description string) error {
	query := "UPDATE tasks SET title = $1, description = $2 WHERE id = $3"
	_, err := database.DB.Exec(context.Background(), query, title, description, id)
	return err
}

// CompleteTask marks a task as completed in the database
func CompleteTask(id int) error {
	query := "UPDATE tasks SET is_completed = TRUE WHERE id = $1"
	_, err := database.DB.Exec(context.Background(), query, id)
	return err
}

// ListTasks retrieves all tasks from the database
func ListTasks() ([]entities.Task, error) {
	query := "SELECT id, title, description, is_completed FROM tasks"
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entities.Task
	for rows.Next() {
		var task entities.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.IsCompleted)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
