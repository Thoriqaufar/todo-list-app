package main

import (
	"github.com/Thoriqaufar/todo-list-app/config"
	category_controller "github.com/Thoriqaufar/todo-list-app/controllers/category-controller"
	completed_task_controller "github.com/Thoriqaufar/todo-list-app/controllers/completed-task-controller"
	index_controller "github.com/Thoriqaufar/todo-list-app/controllers/index-controller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Create Category
	http.HandleFunc("/category/create", category_controller.Create)

	// Active Task
	http.HandleFunc("/", index_controller.Index)
	http.HandleFunc("/active-task", index_controller.Index)
	http.HandleFunc("/tasks/create", index_controller.Create)
	http.HandleFunc("/tasks/detail", index_controller.Detail)
	http.HandleFunc("/tasks/edit", index_controller.Edit)
	http.HandleFunc("/tasks/mark-as-done", index_controller.MarkAsDone)
	http.HandleFunc("/tasks/delete", index_controller.Delete)

	// Completed Task
	http.HandleFunc("/completed-task", completed_task_controller.Index)
	http.HandleFunc("/completed-task/detail", completed_task_controller.Detail)

	log.Println("Server running on port 8080")
	_ = http.ListenAndServe(":8080", nil)
}
