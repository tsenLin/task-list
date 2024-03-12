package models

type Task struct {
	ID int	`json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Status int `json:"status" gorm:"default:0"`
}