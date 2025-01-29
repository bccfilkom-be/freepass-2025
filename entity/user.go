package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID             `json:"userID" gorm:"type:varchar(36);primary_key"`
	Name         string                `json:"name" gorm:"type:varchar(255);not null"`
	Email        string                `json:"email" gorm:"type:varchar(50);unique;not null"`
	Password     string                `json:"password" gorm:"type:varchar(255);not null"`
	Address      string                `json:"address" gorm:"type:varchar(100)"`
	CreatedAt    time.Time             `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    time.Time             `json:"updatedAt" gorm:"autoUpdateTime"`
	RoleID       int                   `json:"roleID"`
	Proposals    []SessionProposal     `json:"proposals" gorm:"foreignKey:UserID"`
	Feedbacks    []Feedback            `json:"feedbacks" gorm:"foreignKey:UserID"`
	Registration []SessionRegistration `json:"registration" gorm:"foreignKey:UserID"`
}
