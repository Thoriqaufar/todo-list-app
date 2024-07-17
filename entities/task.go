package entities

import "time"

type Task struct {
	Id          uint
	Category    Category
	Title       string
	Description string
	Priority    string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
