package main

import (
	"log"

	"github.com/ltphat2204/domain-driven-golang/application"
	"github.com/ltphat2204/domain-driven-golang/config"
	"github.com/ltphat2204/domain-driven-golang/domain"
	"github.com/ltphat2204/domain-driven-golang/handlers"
	"github.com/ltphat2204/domain-driven-golang/infrastructure"
	"github.com/ltphat2204/domain-driven-golang/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	db.AutoMigrate(&domain.Task{})
}

func main() {
	taskRepo := infrastructure.NewTaskRepository(db)
	taskService := application.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := gin.Default()

	routes.SetupRoutes(r, taskHandler)

	r.Run(":8080")
}