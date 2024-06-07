package services

import (
	"errors"
	"ocCourseService/dto"
	"ocCourseService/helper"
	"ocCourseService/models"
	"strconv"

	"gorm.io/gorm"
)

func (s *service) CreateMentor(input dto.MentorInput) (models.Mentor, error) {

	data := models.Mentor{
		Name:       input.Name,
		Profile:    input.Profile,
		Email:      input.Email,
		Profession: input.Profession,
	}

	createdData, err := s.repository.CreateMentor(data)
	if err != nil {

		return data, err
	}

	return createdData, nil
}

func (s *service) UpdateMentor(id string, input dto.UpdateMentorInput) (models.Mentor, error) {
	_, err := s.repository.GetMentorByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Mentor{}, errors.New("mentor not found")
		}
		return models.Mentor{}, err
	}
	data := models.Mentor{
		Name:       input.Name,
		Profile:    input.Profile,
		Email:      input.Email,
		Profession: input.Profession,
	}
	updateMentor, err := s.repository.UpdateMentorByID(id, data)
	if err != nil {
		return updateMentor, err
	}

	return updateMentor, nil
}

func (s *service) GetAllMentor(pageSize string, page string, keyword string) ([]models.Mentor, helper.MetaData, error) {
	convertedPage := 1
	convertedPageSize := 10

	if pageSize != "" {
		formattedPageSize, err := strconv.Atoi(pageSize)
		if err != nil {
			return []models.Mentor{}, helper.MetaData{}, err
		}
		convertedPageSize = formattedPageSize
	}

	if page != "" {
		formattedPage, err := strconv.Atoi(page)
		if err != nil {
			return []models.Mentor{}, helper.MetaData{}, err
		}
		convertedPage = formattedPage
	}

	data, metaData, err := s.repository.GetAllMentor(convertedPageSize, convertedPage, keyword)
	if err != nil {
		return data, metaData, err
	}

	return data, metaData, nil
}

func (s *service) GetMentorByID(id string) (models.Mentor, error) {
	data, err := s.repository.GetMentorByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Mentor{}, errors.New("mentor not found")
		}
		return data, err
	}

	return data, nil
}

func (s *service) DeleteMentorByID(id string) (models.Mentor, error) {
	_, err := s.repository.GetMentorByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Mentor{}, errors.New("mentor not found")
		}
		return models.Mentor{}, err
	}

	data, err := s.repository.DeleteMentorByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}
