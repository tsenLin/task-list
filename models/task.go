package models

type NewTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskRequest struct {
	Name string `json:"name" binding:"required"`
	Status *int `json:"status" binding:"required"`
}