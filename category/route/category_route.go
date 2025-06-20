package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/domain-driven-golang/category/handler"
)

func SetupRoutes(r *gin.Engine, categoryHandler *handler.CategoryHandler) {
	r.POST("/categories", categoryHandler.CreateCategory)
	r.GET("/categories/:id", categoryHandler.GetCategory)
	r.GET("/categories", categoryHandler.GetCategories)
	r.PATCH("/categories/:id", categoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", categoryHandler.DeleteCategory)
}