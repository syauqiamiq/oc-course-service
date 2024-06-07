package models

import "github.com/google/uuid"

type Lesson struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	Name      string
	Video     string
	ChapterID uuid.UUID `gorm:"type:uuid;foreignKey;index"`
	CustomModel
}
