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

func (r *FeedbackRepository) GetAllBySessionId(sessionId uint) ([]domain.SessionFeedback, error) {
	var feedbacks []domain.SessionFeedback
	err := r.db.Preload("User").Where("session_id = ?", sessionId).Find(&feedbacks).Error
	return feedbacks, err
}

func (r *FeedbackRepository) Create(data domain.SessionFeedback) error {
	return r.db.Create(&data).Error
}

func (r *FeedbackRepository) DeleteById(id, sessionId uint) error {
	return r.db.Unscoped().Where("id = ?", id).Where("session_id = ?", sessionId).Delete(&domain.SessionFeedback{}).Error
}

func (r *FeedbackRepository) GetById(id, sessionId uint) (domain.SessionFeedback, error) {
	var feedback domain.SessionFeedback
	err := r.db.Preload("User").Where("id = ?", id).Where("session_id = ?", sessionId).First(&feedback).Error
	return feedback, err
}
