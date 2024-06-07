package models

import "github.com/google/uuid"

type Mentor struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();index;primaryKey"`
	Name       string
	Profile    string
	Email      string
	Profession string
	CustomModel
}
