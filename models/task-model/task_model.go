package task_model

import (
	"github.com/Thoriqaufar/todo-list-app/config"
	"github.com/Thoriqaufar/todo-list-app/entities"
	"github.com/Thoriqaufar/todo-list-app/helper"
	"time"
)

func GetAll() []entities.Task {
	rows, err := config.DB.Query(`
		select
		    tasks.id,
		    tasks.title,
		    categories.name as category_name,
		    tasks.description,
		    tasks.priority,
		    tasks.status,
		    tasks.created_at,
		    tasks.created_at
		from tasks
		join categories on tasks.category_id = categories.id
	`)

	helper.ErrorHandler(err)
	defer rows.Close()

	var tasks []entities.Task

	for rows.Next() {
		var task entities.Task
		err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Category.Name,
			&task.Description,
			&task.Priority,
			&task.Status,
			&task.CreatedAt,
			&task.CreatedAt,
		)

		switch task.Status {
		case "pending":
			task.Status = "Pending"
		case "in_progress":
			task.Status = "In Progress"
		case "completed":
			task.Status = "Completed"
		}

		switch task.Priority {
		case "low":
			task.Priority = "Low"
		case "medium":
			task.Priority = "Medium"
		case "high":
			task.Priority = "High"
		}

		helper.ErrorHandler(err)

		tasks = append(tasks, task)
	}

	return tasks
}

func Create(task entities.Task) bool {
	result, err := config.DB.Exec(`
		insert into tasks(
		                  title, category_id, description, priority, status, created_at, updated_at
		) values (?, ?, ?, ?, ?, ?, ?)`,
		task.Title,
		task.Category.Id,
		task.Description,
		task.Priority,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	)
	helper.ErrorHandler(err)

	lastInsertId, err := result.LastInsertId()
	helper.ErrorHandler(err)

	return lastInsertId > 0
}

func Detail(id int) entities.Task {
	row := config.DB.QueryRow(`
		select
		    tasks.id,
		    tasks.title,
		    categories.name as category_name,
		    tasks.description,
		    tasks.priority,
		    tasks.status,
		    tasks.created_at,
		    tasks.updated_at
		from tasks
		join categories on tasks.category_id = categories.id
		where tasks.id = ?`, id,
	)

	task := entities.Task{}

	err := row.Scan(
		&task.Id,
		&task.Title,
		&task.Category.Name,
		&task.Description,
		&task.Priority,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	helper.ErrorHandler(err)

	switch task.Status {
	case "pending":
		task.Status = "Pending"
	case "in_progress":
		task.Status = "In Progress"
	case "completed":
		task.Status = "Completed"
	}

	switch task.Priority {
	case "low":
		task.Priority = "Low"
	case "medium":
		task.Priority = "Medium"
	case "high":
		task.Priority = "High"
	}

	return task
}

func Update(id int, task entities.Task) bool {
	query, err := config.DB.Exec(`
		update tasks set
		                 title = ?,
		                 category_id = ?,
		                 description = ?,
		                 priority = ?,
		                 status = ?,
		                 updated_at = ?
		where id = ?`,
		task.Title,
		task.Category.Id,
		task.Description,
		task.Priority,
		task.Status,
		task.UpdatedAt,
		id,
	)
	helper.ErrorHandler(err)

	result, err := query.RowsAffected()
	helper.ErrorHandler(err)

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`delete from tasks where id = ?`, id)
	return err
}

func MarkAsDone(id int) error {
	row := config.DB.QueryRow(`
		select
		    tasks.id,
		    tasks.category_id,
		    tasks.title,
		    tasks.description,
		    tasks.priority,
		    tasks.created_at
		from tasks
		where tasks.id = ?`, id,
	)

	compeleted_task := entities.CompletedTask{}

	compeleted_task.CompletedAt = time.Now()
	compeleted_task.Status = "Completed"

	err := row.Scan(
		&compeleted_task.Id,
		&compeleted_task.Category.Id,
		&compeleted_task.Title,
		&compeleted_task.Description,
		&compeleted_task.Priority,
		&compeleted_task.CreatedAt,
	)
	helper.ErrorHandler(err)

	_, err = config.DB.Exec(`
		insert into completed_tasks(
		                  title, category_id, description, priority, status, created_at, completed_at
		) values (?, ?, ?, ?, ?, ?, ?)`,
		compeleted_task.Title,
		compeleted_task.Category.Id,
		compeleted_task.Description,
		compeleted_task.Priority,
		compeleted_task.Status,
		compeleted_task.CreatedAt,
		compeleted_task.CompletedAt,
	)
	helper.ErrorHandler(err)

	_, err = config.DB.Exec(`delete from tasks where id = ?`, id)
	return err
}
