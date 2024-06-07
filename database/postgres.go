package database

import (
	"fmt"
	"log"
	"ocCourseService/models"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	_           = godotenv.Load()
	DB_NAME     = os.Getenv("DB_NAME")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
)

func InitPostgresDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_NAME, DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Unable to connect to database")
		os.Exit(1)
	}

	err = db.AutoMigrate(
		&models.Chapter{},
		&models.Course{},
		&models.ImageCourse{},
		&models.Lesson{},
		&models.Mentor{},
		&models.MyCourse{},
		&models.Review{})
	if err != nil {
		log.Println("Unable to migrate")
		os.Exit(1)
	}

	return db
}
