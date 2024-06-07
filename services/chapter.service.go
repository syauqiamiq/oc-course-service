package services

import (
	"errors"
	"ocCourseService/dto"
	"ocCourseService/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *service) CreateChapter(input dto.ChapterInput) (models.Chapter, error) {

	_, err := s.repository.GetCourseByID(input.CourseID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Chapter{}, errors.New("course not found")
		}
		return models.Chapter{}, err
	}

	data := models.Chapter{
		Name:     input.Name,
		CourseID: uuid.MustParse(input.CourseID),
	}

	createdData, err := s.repository.CreateChapter(data)
	if err != nil {

		return data, err
	}

	return createdData, nil
}

func (s *service) UpdateChapter(id string, input dto.UpdateChapterInput) (models.Chapter, error) {
	_, err := s.repository.GetCourseByID(input.CourseID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Chapter{}, errors.New("course not found")
		}
		return models.Chapter{}, err
	}

	_, err = s.repository.GetChapterByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Chapter{}, errors.New("chapter not found")
		}
		return models.Chapter{}, err
	}

	data := models.Chapter{
		Name:     input.Name,
		CourseID: uuid.MustParse(input.CourseID),
	}
	updateChapter, err := s.repository.UpdateChapterByID(id, data)
	if err != nil {
		return updateChapter, err
	}

	return updateChapter, nil
}

func (s *service) GetAllChapter() ([]models.Chapter, error) {
	data, err := s.repository.GetAllChapter()
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) GetChapterByID(id string) (models.Chapter, error) {
	data, err := s.repository.GetChapterByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Chapter{}, errors.New("chapter not found")
		}
		return data, err
	}

	return data, nil
}

func (s *service) DeleteChapterByID(id string) (models.Chapter, error) {
	_, err := s.repository.GetChapterByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Chapter{}, errors.New("chapter not found")
		}
		return models.Chapter{}, err
	}
	data, err := s.repository.DeleteChapterByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}
