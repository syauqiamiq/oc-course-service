package dto

import (
	"ocCourseService/models"
	"time"
)

type MentorInput struct {
	Name       string `json:"name" binding:"required"`
	Profile    string `json:"profile"`
	Email      string `json:"email" binding:"required"`
	Profession string `json:"profession" binding:"required"`
}

type UpdateMentorInput struct {
	Name       string `json:"name"`
	Profile    string `json:"profile"`
	Email      string `json:"email"`
	Profession string `json:"profession"`
}

type MentorResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Profile    string    `json:"profile"`
	Email      string    `json:"email"`
	Profession string    `json:"profession"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatMentor(data models.Mentor) MentorResponse {

	return MentorResponse{
		ID:         data.ID.String(),
		Name:       data.Name,
		Profile:    data.Profile,
		Email:      data.Email,
		Profession: data.Profession,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}

func FormatListMentor(data []models.Mentor) []MentorResponse {

	formattedData := []MentorResponse{}
	for _, v := range data {
		formattedData = append(formattedData, MentorResponse{
			ID:         v.ID.String(),
			Name:       v.Name,
			Profile:    v.Profile,
			Profession: v.Profession,
			Email:      v.Email,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		})

	}
	return formattedData
}
