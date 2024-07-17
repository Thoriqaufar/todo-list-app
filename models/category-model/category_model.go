package category_model

import (
	"github.com/Thoriqaufar/todo-list-app/config"
	"github.com/Thoriqaufar/todo-list-app/entities"
	"github.com/Thoriqaufar/todo-list-app/helper"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`select * from categories`)
	helper.ErrorHandler(err)

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt)
		helper.ErrorHandler(err)

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
		insert into categories (name, created_at) values (?, ?)`,
		category.Name,
		category.CreatedAt,
	)
	helper.ErrorHandler(err)

	id, err := result.LastInsertId()
	helper.ErrorHandler(err)

	return id > 0
}
