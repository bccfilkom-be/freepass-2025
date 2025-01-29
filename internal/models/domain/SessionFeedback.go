package domain

import (
	"time"

	"gorm.io/gorm"
)

type SessionFeedback struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	SessionID uint    `gorm:"not null"`
	Session   Session `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uint `gorm:"not null"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Content string `gorm:"type:text;not null"`
	Rating  int    `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
