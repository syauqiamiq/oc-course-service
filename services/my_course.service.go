package services

import (
	"errors"
	"fmt"
	"log"
	"ocCourseService/dto"
	"ocCourseService/helper"
	"ocCourseService/models"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)


func (s *service) CreateMyCourse(input dto.MyCourseInput) (models.MyCourse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// CHECK COURSE EXIST
	_, err = s.repository.GetCourseByID(input.CourseID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.MyCourse{}, errors.New("course not found")
		}
		return models.MyCourse{}, err
	}
	data := models.MyCourse{
		UserID:       uuid.MustParse(input.UserID),
		CourseID:    uuid.MustParse(input.CourseID),
	}

	// CHECK USER EXIST
	userResponse, err := helper.ApiRequest("GET", os.Getenv("USER_SERVICE_URL"), fmt.Sprintf("/user/%s", input.UserID), nil)
	if err != nil {
		log.Printf("SERVICE-ERR-CMC 1: %v", err.Error())
		return models.MyCourse{}, err
	}

	if userResponse.Code == 404 && userResponse.Message == "user not found" {
		return models.MyCourse{},errors.New("user not found")
	}

	// CHECK DUPLICATION
	_, err = s.repository.CheckMyCourseIsExist(input.CourseID, input.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			createdData, err := s.repository.CreateMyCourse(data)
			if err != nil {
		
				return data, err
			}
			return createdData, nil
		}
		return models.MyCourse{}, err
	}
	return models.MyCourse{}, errors.New("user already take this course")
}

func (s *service) GetAllMyCourse(userId string) ([]models.MyCourse, error) {
	data, err := s.repository.GetAllMyCourse(userId)
	if err != nil {
		return data, err
	}

	return data, nil
}
