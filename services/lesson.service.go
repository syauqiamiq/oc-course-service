package services

import (
	"errors"
	"ocCourseService/dto"
	"ocCourseService/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *service) CreateLesson(input dto.LessonInput) (models.Lesson, error) {

	_, err := s.repository.GetChapterByID(input.ChapterID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Lesson{}, errors.New("chapter not found")
		}
		return models.Lesson{}, err
	}

	data := models.Lesson{
		Name:      input.Name,
		Video:     input.Video,
		ChapterID: uuid.MustParse(input.ChapterID),
	}

	createdData, err := s.repository.CreateLesson(data)
	if err != nil {

		return data, err
	}

	return createdData, nil
}

func (s *service) UpdateLesson(id string, input dto.UpdateLessonInput) (models.Lesson, error) {
	data := models.Lesson{}

	if input.ChapterID != "" {
		_, err := s.repository.GetChapterByID(input.ChapterID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return models.Lesson{}, errors.New("chapter not found")
			}
			return models.Lesson{}, err
		}
		data = models.Lesson{
			Name:      input.Name,
			Video:     input.Video,
			ChapterID: uuid.MustParse(input.ChapterID),
		}
	} else {
		data = models.Lesson{
			Name:  input.Name,
			Video: input.Video,
		}
	}

	_, err := s.repository.GetLessonByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Lesson{}, errors.New("lesson not found")
		}
		return models.Lesson{}, err
	}

	updateLesson, err := s.repository.UpdateLessonByID(id, data)
	if err != nil {
		return updateLesson, err
	}

	return updateLesson, nil
}

func (s *service) GetAllLesson(chapterId string) ([]models.Lesson, error) {
	data, err := s.repository.GetAllLesson(chapterId)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) GetLessonByID(id string) (models.Lesson, error) {
	data, err := s.repository.GetLessonByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Lesson{}, errors.New("lesson not found")
		}
		return data, err
	}

	return data, nil
}

func (s *service) DeleteLessonByID(id string) (models.Lesson, error) {
	_, err := s.repository.GetLessonByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Lesson{}, errors.New("lesson not found")
		}
		return models.Lesson{}, err
	}
	data, err := s.repository.DeleteLessonByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}
