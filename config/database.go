package config

import (
	"database/sql"
	"github.com/Thoriqaufar/todo-list-app/helper"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/todo_list_db?parseTime=true")
	helper.ErrorHandler(err)

	log.Println("Database Connected")
	DB = db
}
