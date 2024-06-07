package models

import "github.com/google/uuid"

type Course struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	Name        string
	Certificate *bool
	Thumbnail   string
	Type        string
	Status      string
	Price       uint `gorm:"default:0;"`
	Level       string
	Description string
	MentorID    uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	CustomModel
}
