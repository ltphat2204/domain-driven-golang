package routes

import (
	"github.com/ltphat2204/domain-driven-golang/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handler *handlers.TaskHandler) {
	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks/:id", handler.GetTask)
	r.GET("/tasks", handler.GetAllTasks)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
}