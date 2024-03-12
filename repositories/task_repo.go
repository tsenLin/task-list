package repositories

import (
	"task-list/models"

	"gorm.io/gorm"
)

type TaskRepo interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id int) (*models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
	UpdateTask(id int, task *models.Task) (*models.Task, error)
	DeleteTask(id int) (error)
}

type taskRepo struct {
	db *gorm.DB
}

func CreateTaskRepo(db *gorm.DB) TaskRepo{
	// Migrate the schema
	err := db.AutoMigrate(&models.Task{})
	if err != nil {
		panic("failed to migrate database schema")
	}

	return &taskRepo{db: db}
}

func (t *taskRepo) InitTasks() {
	t.db.Create(&models.Task{Name: "Task 1", Status: 0})
	t.db.Create(&models.Task{Name: "Task 2", Status: 1})
}

func (t *taskRepo) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task

	err := t.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *taskRepo) GetTaskById(id int) (*models.Task, error) {
	var task models.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *taskRepo) CreateTask(task *models.Task) (*models.Task, error) {
	err := t.db.Create(&task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepo) UpdateTask(id int, updatedTask *models.Task) (*models.Task, error) {
	updatedTask.ID = id	
	err := t.db.Save(&updatedTask).Error
	if err != nil {
		return nil, err
	}
	return updatedTask, nil
}

func (t *taskRepo) DeleteTask(id int) error {
	err := t.db.Delete(&models.Task{}, id).Error
	if err != nil {
		return err
	}
	return nil
}