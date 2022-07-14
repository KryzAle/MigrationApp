package db

import (
	"time"

	"gorm.io/gorm"
)

type UserOrm struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        string `gorm:"uniqueIndex"`
	Biography    string
	CvUrl        string
	EnglishLevel uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (*UserOrm) TableName() string {
	return "users"
}
