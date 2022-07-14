package db

import (
	"time"

	"gorm.io/gorm"
)

type SkillOrm struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	CategoryId uint
	CreatedAt  time.Time      `gorm:"index"`
	UpdatedAt  time.Time      `gorm:"index"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (SkillOrm) TableName() string {
	return "skills"
}
