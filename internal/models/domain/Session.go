package domain

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`

	RegistrationStartDate time.Time `gorm:"not null;colum:registration_start_date"`
	RegistrationEndDate   time.Time `gorm:"not null;colum:registration_end_date"`

	SesionStartDate time.Time `gorm:"not null;colum:sesion_start_date"`
	SesionEndDate   time.Time `gorm:"not null;colum:sesion_end_date"`

	MaxSeat         int    `gorm:"not null"`
	Status          string `gorm:"type:enum('ACCEPTED','REJECTED', 'PENDING');not null;default:'PENDING'"`
	RejectedMessage string `gorm:"type:varchar(255)"`

	UserID uint `gorm:"not null"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
