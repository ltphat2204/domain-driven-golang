package taskroutes

import (
	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/domain-driven-golang/modules/task/handlers"
)

func SetupRoutes(r *gin.Engine, taskHandler *handlers.TaskHandler) {
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.GET("/tasks", taskHandler.GetTasks)
	r.PATCH("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)
}
