package controllers

import (
	"net/http"
	"task-list/models"
	"task-list/utils"

	"task-list/services"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService services.TaskService
}

type TaskUri struct {
	ID int `uri:"id" binding:"required"`
}

func CreateTaskController(taskService services.TaskService) *taskController {
	return &taskController{
		taskService: taskService,
	}
}

func (t *taskController) GetAllTasks(c *gin.Context) {
    tasks, err := t.taskService.GetAllTasks()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	utils.HandleSuccess(c, http.StatusOK, tasks)
}

func (t *taskController) CreateTask(c *gin.Context) {
    var taskRequest models.NewTaskRequest
	
    if err := c.ShouldBindJSON(&taskRequest); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
        return
    }

	newTask, err := t.taskService.CreateTask(&taskRequest)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	utils.HandleSuccess(c, http.StatusCreated, newTask)
}

func (t *taskController) UpdateTask(c *gin.Context) {
	var taskUri TaskUri
	if err := c.ShouldBindUri(&taskUri); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)		
		return
	}

    var taskRequest models.UpdateTaskRequest

    if err := c.ShouldBindJSON(&taskRequest); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)		
        return
    }

	_, err := t.taskService.GetTaskById(taskUri.ID)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err)
		return
	}

	updatedTask, err := t.taskService.UpdateTask(taskUri.ID, &taskRequest)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	utils.HandleSuccess(c, http.StatusOK, updatedTask)
}

func (t *taskController) DeleteTask(c *gin.Context) {
	var taskUri TaskUri
	if err := c.ShouldBindUri(&taskUri); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	_, err := t.taskService.GetTaskById(taskUri.ID)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err)
		return
	}

	if err := t.taskService.DeleteTask(taskUri.ID); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}