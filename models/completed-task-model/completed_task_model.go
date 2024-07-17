package completed_task_model

import (
	"github.com/Thoriqaufar/todo-list-app/config"
	"github.com/Thoriqaufar/todo-list-app/entities"
	"github.com/Thoriqaufar/todo-list-app/helper"
)

func GetAll() []entities.CompletedTask {
	rows, err := config.DB.Query(`
		select
		    completed_tasks.id,
		    completed_tasks.title,
		    categories.name as category_name,
		    completed_tasks.description,
		    completed_tasks.priority,
		    completed_tasks.status,
		    completed_tasks.created_at,
		    completed_tasks.completed_at
		from completed_tasks
		join categories on completed_tasks.category_id = categories.id
	`)
	helper.ErrorHandler(err)

	defer rows.Close()

	var completedTasks []entities.CompletedTask

	for rows.Next() {
		var completedTask entities.CompletedTask
		err := rows.Scan(
			&completedTask.Id,
			&completedTask.Title,
			&completedTask.Category.Name,
			&completedTask.Description,
			&completedTask.Priority,
			&completedTask.Status,
			&completedTask.CreatedAt,
			&completedTask.CompletedAt,
		)

		switch completedTask.Priority {
		case "low":
			completedTask.Priority = "Low"
		case "medium":
			completedTask.Priority = "Medium"
		case "high":
			completedTask.Priority = "High"
		}

		helper.ErrorHandler(err)

		completedTasks = append(completedTasks, completedTask)
	}

	return completedTasks
}

func Detail(id int) entities.CompletedTask {
	row := config.DB.QueryRow(`
		select
		    completed_tasks.id,
		    completed_tasks.title,
		    categories.name as category_name,
		    completed_tasks.description,
		    completed_tasks.priority,
		    completed_tasks.status,
		    completed_tasks.created_at,
		    completed_tasks.completed_at
		from completed_tasks
		join categories on completed_tasks.category_id = categories.id
		where completed_tasks.id = ?`, id,
	)

	completedTask := entities.CompletedTask{}

	err := row.Scan(
		&completedTask.Id,
		&completedTask.Title,
		&completedTask.Category.Name,
		&completedTask.Description,
		&completedTask.Priority,
		&completedTask.Status,
		&completedTask.CreatedAt,
		&completedTask.CompletedAt,
	)
	helper.ErrorHandler(err)

	switch completedTask.Status {
	case "pending":
		completedTask.Status = "Pending"
	case "in_progress":
		completedTask.Status = "In Progress"
	case "completed":
		completedTask.Status = "Completed"
	}

	switch completedTask.Priority {
	case "low":
		completedTask.Priority = "Low"
	case "medium":
		completedTask.Priority = "Medium"
	case "high":
		completedTask.Priority = "High"
	}

	return completedTask
}
