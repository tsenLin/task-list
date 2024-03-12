package models

type TaskRequest struct {
	Name string `json:"name" binding:"required"`
	Status *int `json:"status" binding:"required"`
}