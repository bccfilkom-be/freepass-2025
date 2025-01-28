package repository

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/services/feedback"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) feedback.FeedbackRepository {
	return &FeedbackRepository{db}
}

func (r *FeedbackRepository) Create(data domain.SessionFeedback) error {
	return r.db.Create(&data).Error
}
