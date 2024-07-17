package category_controller

import (
	"github.com/Thoriqaufar/todo-list-app/entities"
	"github.com/Thoriqaufar/todo-list-app/helper"
	category_model "github.com/Thoriqaufar/todo-list-app/models/category-model"
	"html/template"
	"net/http"
	"time"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.gohtml")
		helper.ErrorHandler(err)

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()

		ok := category_model.Create(category)
		if !ok {
			temp, err := template.ParseFiles("views/category/create.html")
			helper.ErrorHandler(err)

			_ = temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/tasks/create", http.StatusSeeOther)
	}
}
