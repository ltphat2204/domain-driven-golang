package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	categoryApplication "github.com/ltphat2204/domain-driven-golang/category/application"
	categoryDomain "github.com/ltphat2204/domain-driven-golang/category/domain"
	categoryHandler "github.com/ltphat2204/domain-driven-golang/category/handler"
	categoryInfrastructure "github.com/ltphat2204/domain-driven-golang/category/infrastructure"
	categoryRoutes "github.com/ltphat2204/domain-driven-golang/category/route"
	"github.com/ltphat2204/domain-driven-golang/config"
	"github.com/ltphat2204/domain-driven-golang/routes"
	taskApplication "github.com/ltphat2204/domain-driven-golang/task/application"
	taskDomain "github.com/ltphat2204/domain-driven-golang/task/domain"
	taskHandler "github.com/ltphat2204/domain-driven-golang/task/handlers"
	taskInfrastructure "github.com/ltphat2204/domain-driven-golang/task/infrastructure"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, using default environment variables instead")
	}

	db, err = config.GetDb()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&taskDomain.Task{}, &categoryDomain.Category{})
}

func main() {
	taskRepo := taskInfrastructure.NewTaskRepository(db)
	taskService := taskApplication.NewTaskService(taskRepo)
	taskHandler := taskHandler.NewTaskHandler(taskService)
	
	categoryRepo := categoryInfrastructure.NewCategoryRepository(db)
	categoryService := categoryApplication.NewCategoryService(categoryRepo)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryService)

	r := gin.Default()

	categoryRoutes.SetupRoutes(r, categoryHandler)
	routes.SetupRoutes(r, taskHandler)

	r.Run(":8080")
}