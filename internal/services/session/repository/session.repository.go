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

func (v *SessionRepository) GetAll(filter session.SessionFilter) ([]domain.Session, error) {
	var data []domain.Session
	query := v.db.Model(&domain.Session{})
	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	err := query.Preload("User").Find(&data).Error
	return data, err
}
