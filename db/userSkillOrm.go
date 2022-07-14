package db

import (
	"time"

	"gorm.io/gorm"
)

type UserSkillOrm struct {
	ID uint `gorm:"primaryKey"`
	// We are using UserEmail instead of UserId because Users are stored in Notion and we reference users by Email
	UserEmail  string `gorm:"uniqueIndex:idx_useremail_skillid"`
	SkillID    uint   `gorm:"uniqueIndex:idx_useremail_skillid"`
	Experience uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (UserSkillOrm) TableName() string {
	return "user_skills"
}
