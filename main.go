package main

import (
	"fmt"
	"log"
	"ocCourseService/database"
	"ocCourseService/handler"
	"ocCourseService/repository"
	"ocCourseService/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.InitPostgresDB()

	repository := repository.NewRepository(db)

	service := services.NewService(repository)

	_ = handler.NewHandler(service)

	router := gin.Default()

	// courseRouter := router.Group("/course")

	err = router.Run(fmt.Sprintf(":%s", os.Getenv("RUNNING_PORT")))
	if err != nil {
		panic("Error When Running")
	}
}
