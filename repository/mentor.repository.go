package repository

import (
	"math"
	"ocCourseService/helper"
	"ocCourseService/models"
	"strings"
)

func (r *repository) CreateMentor(data models.Mentor) (models.Mentor, error) {

	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetAllMentor(pageSize int, page int, keyword string) ([]models.Mentor, helper.MetaData, error) {

	var metaData helper.MetaData
	var data []models.Mentor

	db := r.db.Model(&data)

	if keyword != "" {
		db.Where("LOWER(name)  LIKE ?", "%"+strings.ToLower(keyword)+"%")
	}

	// Calculate the total number of records
	var totalRecords int64
	db.Count(&totalRecords)

	// Calculate the last_page based on totalRecords and pageSize
	lastPage := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	// Calculate the offset based on page number and page size
	offset := (page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize)

	err := db.Scan(&data).Error
	if err != nil {
		return data, metaData, err
	}
	metaData.CurrentPage = page
	metaData.LastPage = lastPage
	metaData.PageSize = pageSize
	metaData.Total = int(totalRecords)

	return data, metaData, nil

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
