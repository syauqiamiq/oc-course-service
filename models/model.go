package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
