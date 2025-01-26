package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Bio       string `gorm:"type:text;not null"`
	Role      string `gorm:"type:enum('ADMIN', 'COORDINATOR', 'USER');not null;default:'USER'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
