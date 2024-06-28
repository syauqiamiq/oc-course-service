package services

import (
	"errors"
	"ocCourseService/dto"
	"ocCourseService/helper"
	"ocCourseService/models"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *service) CreateCourse(input dto.CourseInput) (models.Course, error) {

	_, err := s.repository.GetMentorByID(input.MentorID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Course{}, errors.New("mentor not found")
		}
		return models.Course{}, err
	}

	data := models.Course{
		Name:        input.Name,
		Certificate: input.Certificate,
		Thumbnail:   input.Thumbnail,
		Type:        input.Type,
		Status:      input.Status,
		Price:       input.Price,
		Level:       input.Level,
		Description: input.Description,
		MentorID:    uuid.MustParse(input.MentorID),
	}

	createdData, err := s.repository.CreateCourse(data)
	if err != nil {

		return data, err
	}

	return createdData, nil
}

func (s *service) UpdateCourse(id string, input dto.UpdateCourseInput) (models.Course, error) {
	data := models.Course{}

	if input.MentorID != "" {
		_, err := s.repository.GetMentorByID(input.MentorID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return models.Course{}, errors.New("mentor not found")
			}
			return models.Course{}, err
		}

		data = models.Course{
			Name:        input.Name,
			Certificate: input.Certificate,
			Thumbnail:   input.Thumbnail,
			Type:        input.Type,
			Status:      input.Status,
			Price:       input.Price,
			Level:       input.Level,
			Description: input.Description,
			MentorID:    uuid.MustParse(input.MentorID),
		}
	} else {
		data = models.Course{
			Name:        input.Name,
			Certificate: input.Certificate,
			Thumbnail:   input.Thumbnail,
			Type:        input.Type,
			Status:      input.Status,
			Price:       input.Price,
			Level:       input.Level,
			Description: input.Description,
		}
	}

	_, err := s.repository.GetCourseByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Course{}, errors.New("course not found")
		}
		return models.Course{}, err
	}

	updateCourse, err := s.repository.UpdateCourseByID(id, data)
	if err != nil {
		return updateCourse, err
	}

	return updateCourse, nil
}

func (s *service) GetAllCourse(pageSize string, page string, keyword string) ([]models.Course, helper.MetaData, error) {
	convertedPage := 1
	convertedPageSize := 10

	if pageSize != "" {
		formattedPageSize, err := strconv.Atoi(pageSize)
		if err != nil {
			return []models.Course{}, helper.MetaData{}, err
		}
		convertedPageSize = formattedPageSize
	}

	if page != "" {
		formattedPage, err := strconv.Atoi(page)
		if err != nil {
			return []models.Course{}, helper.MetaData{}, err
		}
		convertedPage = formattedPage
	}

	data, metaData, err := s.repository.GetAllCourse(convertedPageSize, convertedPage, keyword)
	if err != nil {
		return data, metaData, err
	}

	return data, metaData, nil
}

func (s *service) GetCourseByID(id string) (models.Course, error) {
	data, err := s.repository.GetCourseByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Course{}, errors.New("course not found")
		}
		return data, err
	}

	return data, nil
}

func (s *service) DeleteCourseByID(id string) (models.Course, error) {
	_, err := s.repository.GetCourseByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Course{}, errors.New("course not found")
		}
		return models.Course{}, err
	}
	data, err := s.repository.DeleteCourseByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}
