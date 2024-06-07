package dto

import (
	"ocCourseService/models"
)

type ChapterInput struct {
	Name     string `json:"name" binding:"required"`
	CourseID string `json:"course_id"  binding:"required"`
}

type UpdateChapterInput struct {
	Name     string `json:"name"`
	CourseID string `json:"course_id"`
}

type ChapterResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	CourseID string `json:"course_id"`
}

func FormatChapter(data models.Chapter) ChapterResponse {

	return ChapterResponse{
		ID:       data.ID.String(),
		Name:     data.Name,
		CourseID: data.CourseID.String(),
	}
}

func FormatListChapter(data []models.Chapter) []ChapterResponse {

	formattedData := []ChapterResponse{}
	for _, v := range data {
		formattedData = append(formattedData, ChapterResponse{
			ID:       v.ID.String(),
			Name:     v.Name,
			CourseID: v.CourseID.String(),
		})

	}
	return formattedData
}
