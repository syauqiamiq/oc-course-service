package repository

import "ocCourseService/models"

func (r *repository) CreateChapter(data models.Chapter) (models.Chapter, error) {

	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetAllChapter(courseId string) ([]models.Chapter, error) {

	var data []models.Chapter

	db := r.db.Model(&data)

	if courseId != "" {
		db.Where("course_id = ?", courseId)
	}

	err := db.Scan(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *repository) GetChapterByID(id string) (models.Chapter, error) {
	var result models.Chapter
	err := r.db.First(&result, "id = ?", id).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) UpdateChapterByID(id string, data models.Chapter) (models.Chapter, error) {
	var result models.Chapter
	err := r.db.Model(&result).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) DeleteChapterByID(id string) (models.Chapter, error) {

	var data models.Chapter
	err := r.db.Delete(&data, "id = ?", id).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
