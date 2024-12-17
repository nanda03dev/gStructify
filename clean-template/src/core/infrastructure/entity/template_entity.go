package entity

import (
	"time"

	"gorm.io/gorm"
)

type TemplateEntity struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
