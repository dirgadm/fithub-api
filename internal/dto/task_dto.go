package dto

import "time"

type TaskResponse struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	User        UserResponse `json:"user,omitempty"`
	DueDate     time.Time    `json:"created_at,omitempty"`
}

type CreateTaskRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required,gte=8"`
	Duedate     time.Time `json:"due_date" validate:"required"`
}
