package services

import (
	"ocCourseService/dto"
	"ocCourseService/helper"
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
	GetAllMentor(pageSize string, page string, keyword string) ([]models.Mentor, helper.MetaData, error)
	GetMentorByID(id string) (models.Mentor, error)
	DeleteMentorByID(id string) (models.Mentor, error)
	CreateCourse(input dto.CourseInput) (models.Course, error)
	UpdateCourse(id string, input dto.UpdateCourseInput) (models.Course, error)
	GetAllCourse(pageSize string, page string, keyword string) ([]models.Course, helper.MetaData, error)
	GetCourseByID(id string) (models.Course, error)
	DeleteCourseByID(id string) (models.Course, error)
	CreateChapter(input dto.ChapterInput) (models.Chapter, error)
	UpdateChapter(id string, input dto.UpdateChapterInput) (models.Chapter, error)
	GetAllChapter() ([]models.Chapter, error)
	GetChapterByID(id string) (models.Chapter, error)
	DeleteChapterByID(id string) (models.Chapter, error)
}
