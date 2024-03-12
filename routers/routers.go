package routers

import (
	"net/http"
	"task-list/config"
	"task-list/controllers"
	"task-list/middlewares/auth"
	"task-list/repositories"
	"task-list/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	

	taskRepo := repositories.CreateTaskRepo(config.DbConnect())
	taskService := services.CreateTaskService(taskRepo)
	taskController := controllers.CreateTaskController(taskService)

	// Group routes
	api := r.Group("/v1")

	// Routes for tasks
	api.GET("/tasks", taskController.GetAllTasks)

	// Routes for task
	taskRoutes := api.Group("/task").Use(auth.ValidAuth)
	{
		taskRoutes.POST("", taskController.CreateTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}

	return r
}