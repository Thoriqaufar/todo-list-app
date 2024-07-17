package entities

import "time"

type CompletedTask struct {
	Id          uint
	Category    Category
	Title       string
	Description string
	Priority    string
	Status      string
	CreatedAt   time.Time
	CompletedAt time.Time
}
