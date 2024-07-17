package completed_task_controller

import (
	"github.com/Thoriqaufar/todo-list-app/helper"
	completed_task_model "github.com/Thoriqaufar/todo-list-app/models/completed-task-model"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	completedTasks := completed_task_model.GetAll()
	data := map[string]any{
		"completedTasks": completedTasks,
	}

	temp, err := template.ParseFiles("views/completed-task/index.gohtml")
	helper.ErrorHandler(err)

	_ = temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	helper.ErrorHandler(err)

	completedTask := completed_task_model.Detail(id)
	data := map[string]any{
		"completedTask": completedTask,
	}

	temp, err := template.ParseFiles("views/completed-task/detail.gohtml")
	helper.ErrorHandler(err)

	_ = temp.Execute(w, data)
}
