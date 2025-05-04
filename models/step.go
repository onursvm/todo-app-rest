package models

import "time"

type Step struct {
	ID        string     `json:"id"`
	ToDoID    string     `json:"todo_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Content   string     `json:"content"`
	Done      bool       `json:"done"`
}
