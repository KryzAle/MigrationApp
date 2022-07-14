package db

import (
	"time"

	"gorm.io/gorm"
)

type UserSkillOrm struct {
	ID         uint   `gorm:"primaryKey"`
	UserId     uint   `gorm:"uniqueIndex:idx_userid_skillid"`
	SkillID    uint   `gorm:"uniqueIndex:idx_userid_skillid"`
	UserEmail  string `gorm:"index,unique"`
	Experience uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (UserSkillOrm) TableName() string {
	return "user_skills"
}
