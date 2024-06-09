package dto

import (
	"ocCourseService/models"
	"time"
)

type MyCourseInput struct {
	UserID       string `json:"user_id" binding:"required"`
	CourseID    string `json:"course_id"  binding:"required"`
}

type MyCourseResponse struct {
	ID         string    `json:"id"`
	UserID       string `json:"user_id"`
	CourseID    string `json:"course_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatMyCourse(data models.MyCourse) MyCourseResponse {

	return MyCourseResponse{
		ID:         data.ID.String(),
		CourseID: data.CourseID.String(),
		UserID: data.UserID.String(),
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func FormatListMyCourse(data []models.MyCourse) []MyCourseResponse {

	formattedData := []MyCourseResponse{}
	for _, v := range data {
		formattedData = append(formattedData, MyCourseResponse{
			ID:         v.ID.String(),
			CourseID: v.CourseID.String(),
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		})

	}
	return formattedData
}
