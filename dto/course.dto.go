package dto

import (
	"ocCourseService/models"
	"time"
)

type CourseInput struct {
	Name        string `json:"name" binding:"required"`
	Certificate *bool  `json:"certificate" binding:"required"`
	Thumbnail   string `json:"thumbnail"`
	Type        string `json:"type" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Price       uint   `json:"price"`
	Level       string `json:"level" binding:"required"`
	Description string `json:"description" `
	MentorID    string `json:"mentor_id" binding:"required"`
}

type UpdateCourseInput struct {
	Name        string `json:"name"`
	Certificate *bool  `json:"certificate"`
	Thumbnail   string `json:"thumbnail"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Price       uint   `json:"price"`
	Level       string `json:"level"`
	Description string `json:"description" `
	MentorID    string `json:"mentor_id"`
}

type CourseResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Certificate *bool     `json:"certificate"`
	Thumbnail   string    `json:"thumbnail"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Price       uint      `json:"price"`
	Level       string    `json:"level"`
	Description string    `json:"description" `
	MentorID    string    `json:"mentor_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatCourse(data models.Course) CourseResponse {

	return CourseResponse{
		ID:          data.ID.String(),
		Name:        data.Name,
		Certificate: data.Certificate,
		Thumbnail:   data.Thumbnail,
		Type:        data.Type,
		Status:      data.Status,
		Price:       data.Price,
		Level:       data.Level,
		Description: data.Description,
		MentorID:    data.MentorID.String(),
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func FormatListCourse(data []models.Course) []CourseResponse {

	formattedData := []CourseResponse{}
	for _, v := range data {
		formattedData = append(formattedData, CourseResponse{
			ID:          v.ID.String(),
			Name:        v.Name,
			Certificate: v.Certificate,
			Thumbnail:   v.Thumbnail,
			Type:        v.Type,
			Status:      v.Status,
			Price:       v.Price,
			Level:       v.Level,
			Description: v.Description,
			MentorID:    v.MentorID.String(),
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})

	}
	return formattedData
}
