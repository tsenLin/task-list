package services

import (
	"task-list/models"
	"task-list/repositories"
)


type TaskService interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id int) (*models.Task, error)
	CreateTask(task *models.NewTaskRequest) (*models.Task, error)
	UpdateTask(id int, task *models.UpdateTaskRequest) (*models.Task, error)
	DeleteTask(id int) (error)
}

type taskService struct {
	taskRepo repositories.TaskRepo
}

func CreateTaskService(taskRepo repositories.TaskRepo) TaskService{
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (t *taskService) GetAllTasks() ([]models.Task, error) {
	return t.taskRepo.GetAllTasks()
}

func (t *taskService) GetTaskById(id int) (*models.Task, error) {
	return t.taskRepo.GetTaskById(id)
}

func (t *taskService) CreateTask(newTask *models.NewTaskRequest) (*models.Task, error) {
	task := models.Task{
		Name: newTask.Name,
	}
	return t.taskRepo.CreateTask(&task)
}

func (t *taskService) UpdateTask(id int, updatedTask *models.UpdateTaskRequest) (*models.Task, error) {
	task := models.Task{
		Name: updatedTask.Name,
		Status: *updatedTask.Status,
	}
	return t.taskRepo.UpdateTask(id, &task)
}

func (t *taskService) DeleteTask(id int) error {
	return t.taskRepo.DeleteTask(id)
}
