package services

import (
	"ocCourseService/dto"
	"ocCourseService/models"
	"ocCourseService/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {

	return &service{
		repository: repository,
	}
}

type Service interface {
	CreateMentor(input dto.MentorInput) (models.Mentor, error)
	UpdateMentor(id string, input dto.UpdateMentorInput) (models.Mentor, error)
	GetAllMentor() ([]models.Mentor, error)
	GetMentorByID(id string) (models.Mentor, error)
	DeleteMentorByID(id string) (models.Mentor, error)
}
