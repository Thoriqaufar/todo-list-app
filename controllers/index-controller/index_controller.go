package index_controller

import (
	"github.com/Thoriqaufar/todo-list-app/entities"
	"github.com/Thoriqaufar/todo-list-app/helper"
	category_model "github.com/Thoriqaufar/todo-list-app/models/category-model"
	task_model "github.com/Thoriqaufar/todo-list-app/models/task-model"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tasks := task_model.GetAll()
	data := map[string]any{
		"tasks": tasks,
	}

	temp, err := template.ParseFiles("views/active-task/index.gohtml")
	helper.ErrorHandler(err)

	_ = temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/active-task/create.gohtml")
		helper.ErrorHandler(err)

		categories := category_model.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		_ = temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var task entities.Task

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		helper.ErrorHandler(err)

		task.Title = r.FormValue("title")
		task.Category.Id = uint(categoryId)
		task.Description = r.FormValue("description")
		task.Priority = r.FormValue("priority")
		task.Status = r.FormValue("status")
		task.CreatedAt = time.Now()
		task.UpdatedAt = time.Now()

		ok := task_model.Create(task)
		if !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	helper.ErrorHandler(err)

	task := task_model.Detail(id)
	data := map[string]any{
		"task": task,
	}

	temp, err := template.ParseFiles("views/active-task/detail.gohtml")
	helper.ErrorHandler(err)

	_ = temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/active-task/edit.gohtml")
		helper.ErrorHandler(err)

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		helper.ErrorHandler(err)

		task := task_model.Detail(id)

		categories := category_model.GetAll()
		data := map[string]any{
			"categories": categories,
			"task":       task,
		}

		_ = temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var task entities.Task

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		helper.ErrorHandler(err)

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		helper.ErrorHandler(err)

		task.Title = r.FormValue("title")
		task.Category.Id = uint(categoryId)
		task.Description = r.FormValue("description")
		task.Priority = r.FormValue("priority")
		task.Status = r.FormValue("status")
		task.UpdatedAt = time.Now()

		ok := task_model.Update(id, task)
		if !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func MarkAsDone(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	helper.ErrorHandler(err)

	err = task_model.MarkAsDone(id)
	helper.ErrorHandler(err)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	helper.ErrorHandler(err)

	err = task_model.Delete(id)
	helper.ErrorHandler(err)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
