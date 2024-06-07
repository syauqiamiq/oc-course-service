package models

import "github.com/google/uuid"

type MyCourse struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	CourseID uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	UserID   uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	CustomModel
}
