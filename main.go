package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	categoryApplication "github.com/ltphat2204/domain-driven-golang/modules/category/application"
	categoryDomain "github.com/ltphat2204/domain-driven-golang/modules/category/domain"
	categoryHandler "github.com/ltphat2204/domain-driven-golang/modules/category/handler"
	categoryInfrastructure "github.com/ltphat2204/domain-driven-golang/modules/category/infrastructure"
	categoryRoutes "github.com/ltphat2204/domain-driven-golang/modules/category/route"

	taskApplication "github.com/ltphat2204/domain-driven-golang/modules/task/application"
	taskDomain "github.com/ltphat2204/domain-driven-golang/modules/task/domain"
	taskHandler "github.com/ltphat2204/domain-driven-golang/modules/task/handlers"
	taskInfrastructure "github.com/ltphat2204/domain-driven-golang/modules/task/infrastructure"
	taskRoutes "github.com/ltphat2204/domain-driven-golang/modules/task/route"

	"github.com/ltphat2204/domain-driven-golang/config"
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
	taskRoutes.SetupRoutes(r, taskHandler)

	r.Run(":8080")
}
