package repository

import "ocCourseService/models"

func (r *repository) CreateLesson(data models.Lesson) (models.Lesson, error) {

	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetAllLesson(chapterId string) ([]models.Lesson, error) {

	var data []models.Lesson

	db := r.db.Model(&data)

	if chapterId != "" {
		db.Where("chapter_id = ?", chapterId)
	}

	err := db.Scan(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *repository) GetLessonByID(id string) (models.Lesson, error) {
	var result models.Lesson
	err := r.db.First(&result, "id = ?", id).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) UpdateLessonByID(id string, data models.Lesson) (models.Lesson, error) {
	var result models.Lesson
	err := r.db.Model(&result).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) DeleteLessonByID(id string) (models.Lesson, error) {

	var data models.Lesson
	err := r.db.Delete(&data, "id = ?", id).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
