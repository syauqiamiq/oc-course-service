package repository

import "ocCourseService/models"

func (r *repository) CreateMyCourse(data models.MyCourse) (models.MyCourse, error) {

	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}


func (r *repository) CheckMyCourseIsExist(courseId string, userId string) (models.MyCourse, error) {
	var result models.MyCourse
	err := r.db.First(&result, "user_id = ? AND course_id = ?", userId, courseId).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) GetAllMyCourse(userId string) ([]models.MyCourse, error) {

	var data []models.MyCourse

	db := r.db.Model(&data)

	if userId != "" {
		db.Where("user_id = ?", userId)
	}

	err := db.Scan(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}