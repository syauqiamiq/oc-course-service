package models

import "github.com/google/uuid"

type Course struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	Name        string
	Certificate int
	Thumbnail   string
	Type        string
	Status      string
	Price       uint
	Level       string
	Description string
	MentorID    uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	CustomModel
}
