package repository

import (
	"ocCourseService/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{
		db: db,
	}
}

type Repository interface {
	CreateMentor(data models.Mentor) (models.Mentor, error)
	GetAllMentor() ([]models.Mentor, error)
	GetMentorByID(id string) (models.Mentor, error)
	UpdateMentorByID(id string, data models.Mentor) (models.Mentor, error)
	DeleteMentorByID(id string) (models.Mentor, error)
}
