package models

import "github.com/google/uuid"

type Review struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	Note     uint
	Rating   int
	CourseID uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	UserID   uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	CustomModel
}
