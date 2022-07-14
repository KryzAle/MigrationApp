package db

import (
	"time"

	"gorm.io/gorm"
)

type CategoryOrm struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (CategoryOrm) TableName() string {
	return "categories"
}
