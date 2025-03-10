package task

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      *uint  `json:"user_id"`
}

type TaskResponse struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      *uint  `json:"user_id"`
}

func NewTaskResponse(t Task) TaskResponse {
	return TaskResponse{
		Id:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		UserID:      t.UserID,
	}
}

type CreateBody struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Completed   bool   `json:"completed" validate:"boolean"`
	UserID      uint   `json:"user_id" validate:"required"`
}

type PatchBody struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}
