package models

import "github.com/google/uuid"

type Chapter struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	Name     string
	CourseID uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	CustomModel
}
