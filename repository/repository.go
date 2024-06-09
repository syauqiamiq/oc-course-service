package repository

import (
	"ocCourseService/helper"
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
	GetAllMentor(pageSize int, page int, keyword string) ([]models.Mentor, helper.MetaData, error)
	GetMentorByID(id string) (models.Mentor, error)
	UpdateMentorByID(id string, data models.Mentor) (models.Mentor, error)
	DeleteMentorByID(id string) (models.Mentor, error)
	CreateCourse(data models.Course) (models.Course, error)
	GetAllCourse(pageSize int, page int, keyword string) ([]models.Course, helper.MetaData, error)
	GetCourseByID(id string) (models.Course, error)
	UpdateCourseByID(id string, data models.Course) (models.Course, error)
	DeleteCourseByID(id string) (models.Course, error)
	CreateChapter(data models.Chapter) (models.Chapter, error)
	GetAllChapter(courseId string) ([]models.Chapter, error)
	GetChapterByID(id string) (models.Chapter, error)
	UpdateChapterByID(id string, data models.Chapter) (models.Chapter, error)
	DeleteChapterByID(id string) (models.Chapter, error)
	CreateLesson(data models.Lesson) (models.Lesson, error)
	GetAllLesson(chapterId string) ([]models.Lesson, error)
	GetLessonByID(id string) (models.Lesson, error)
	UpdateLessonByID(id string, data models.Lesson) (models.Lesson, error)
	DeleteLessonByID(id string) (models.Lesson, error)
	CreateMyCourse(data models.MyCourse) (models.MyCourse, error)
	CheckMyCourseIsExist(courseId string, userId string)(models.MyCourse, error)
	GetAllMyCourse(userId string) ([]models.MyCourse, error)
}
