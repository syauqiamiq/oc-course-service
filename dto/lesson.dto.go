package dto

import (
	"ocCourseService/models"
)

type LessonInput struct {
	Name      string `json:"name" binding:"required"`
	Video     string `json:"video" binding:"required"`
	ChapterID string `json:"chapter_id"  binding:"required"`
}

type UpdateLessonInput struct {
	Name      string `json:"name" binding:"required"`
	Video     string `json:"video" binding:"required"`
	ChapterID string `json:"chapter_id"`
}

type LessonResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Video     string `json:"video"`
	ChapterID string `json:"chapter_id"`
}

func FormatLesson(data models.Lesson) LessonResponse {

	return LessonResponse{
		ID:        data.ID.String(),
		Name:      data.Name,
		Video:     data.Video,
		ChapterID: data.ChapterID.String(),
	}
}

func FormatListLesson(data []models.Lesson) []LessonResponse {

	formattedData := []LessonResponse{}
	for _, v := range data {
		formattedData = append(formattedData, LessonResponse{
			ID:        v.ID.String(),
			Name:      v.Name,
			Video:     v.Video,
			ChapterID: v.ChapterID.String(),
		})

	}
	return formattedData
}
