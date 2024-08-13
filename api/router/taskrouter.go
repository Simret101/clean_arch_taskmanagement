package router

import (
	"task/api/controller"
	"task/api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(r *gin.Engine, taskController *controller.TaskController, authMiddleware *middleware.AuthMiddleware) {
	// Protected routes
	authorized := r.Group("/")
	authorized.Use(authMiddleware.AuthMiddleware())
	{
		authorized.GET("/tasks", taskController.GetAllTasks)
		authorized.GET("/tasks/:id", taskController.GetTaskByID)
		authorized.POST("/tasks", taskController.CreateTask)
		authorized.PUT("/tasks/:id", taskController.UpdateTask)
		authorized.DELETE("/tasks/:id", taskController.DeleteTask)
	}

	admin := r.Group("/admin")
	admin.Use(authMiddleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/tasks", taskController.GetAllTasks)
		admin.GET("/tasks/:id", taskController.GetTaskByID)
		admin.POST("/tasks", taskController.CreateTask)
		admin.PUT("/tasks/:id", taskController.UpdateTask)
		admin.DELETE("/tasks/:id", taskController.DeleteTask)
	}
}
