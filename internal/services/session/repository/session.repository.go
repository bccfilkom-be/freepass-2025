package repository

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) session.SessionRepository {
	return &SessionRepository{db}
}

func (v *SessionRepository) Create(data domain.Session) error {
	return v.db.Create(&data).Error
}
