package models

import "github.com/google/uuid"

type ImageCourse struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	Image    string
	CourseID uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	CustomModel
}
