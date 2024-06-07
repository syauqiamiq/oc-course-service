package repository

import "ocCourseService/models"

func (r *repository) CreateMentor(data models.Mentor) (models.Mentor, error) {

	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetAllMentor() ([]models.Mentor, error) {
	var result []models.Mentor
	err := r.db.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) GetMentorByID(id string) (models.Mentor, error) {
	var result models.Mentor
	err := r.db.First(&result, "id = ?", id).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) UpdateMentorByID(id string, data models.Mentor) (models.Mentor, error) {
	var result models.Mentor
	err := r.db.Model(&result).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) DeleteMentorByID(id string) (models.Mentor, error) {

	var data models.Mentor
	err := r.db.Delete(&data, "id = ?", id).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
