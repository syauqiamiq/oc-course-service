package services

import (
	"ocCourseService/dto"
	"ocCourseService/models"
)

func (s *service) CreateMentor(input dto.MentorInput) (models.Mentor, error) {

	data := models.Mentor{
		Name:       input.Name,
		Profile:    input.Profile,
		Email:      input.Email,
		Profession: input.Profession,
	}

	createdData, err := s.repository.CreateMentor(data)
	if err != nil {

		return data, err
	}

	return createdData, nil
}

func (s *service) UpdateMentor(id string, input dto.UpdateMentorInput) (models.Mentor, error) {

	data := models.Mentor{
		Name:       input.Name,
		Profile:    input.Profile,
		Email:      input.Email,
		Profession: input.Profession,
	}
	updateMentor, err := s.repository.UpdateMentorByID(id, data)
	if err != nil {
		return updateMentor, err
	}

	return updateMentor, nil
}

func (s *service) GetAllMentor() ([]models.Mentor, error) {
	data, err := s.repository.GetAllMentor()
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) GetMentorByID(id string) (models.Mentor, error) {
	data, err := s.repository.GetMentorByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (s *service) DeleteMentorByID(id string) (models.Mentor, error) {
	data, err := s.repository.DeleteMentorByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}
