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

	handler := handler.NewHandler(service)

	router := gin.Default()

	mentorRouter := router.Group("/mentor")

	mentorRouter.POST("/", handler.CreateMentorHandler)
	mentorRouter.PUT("/:id", handler.UpdateMentorHandler)
	mentorRouter.GET("/:id", handler.GetMentorByIDHandler)
	mentorRouter.DELETE("/:id", handler.DeleteMentorByIDHandler)
	mentorRouter.GET("/", handler.GetMentorHandler)

	courseRouter := router.Group("/course")

	courseRouter.POST("/", handler.CreateCourseHandler)
	courseRouter.PUT("/:id", handler.UpdateCourseHandler)
	courseRouter.GET("/:id", handler.GetCourseByIDHandler)
	courseRouter.DELETE("/:id", handler.DeleteCourseByIDHandler)
	courseRouter.GET("/", handler.GetCourseHandler)

	chapterRouter := router.Group("/chapter")

	chapterRouter.POST("/", handler.CreateChapterHandler)
	chapterRouter.PUT("/:id", handler.UpdateChapterHandler)
	chapterRouter.GET("/:id", handler.GetChapterByIDHandler)
	chapterRouter.DELETE("/:id", handler.DeleteChapterByIDHandler)
	chapterRouter.GET("/", handler.GetChapterHandler)

	err = router.Run(fmt.Sprintf(":%s", os.Getenv("RUNNING_PORT")))
	if err != nil {
		panic("Error When Running")
	}
}
