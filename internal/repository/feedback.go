package repository

import (
	"freepass-bcc/entity"

	"gorm.io/gorm"
)

type IFeedbackRepository interface {
	AddFeedback(feedback *entity.Feedback) error
	GetFeedbackBySessionID(sessionID int) ([]*entity.Feedback, error)
	DeleteFeedbackInSession(feedbackID int) error
}

type FeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) IFeedbackRepository {
	return &FeedbackRepository{db}
}

func (fr *FeedbackRepository) AddFeedback(feedback *entity.Feedback) error {
	err := fr.db.Debug().Create(feedback).Error
	if err != nil {
		return err
	}

	return nil
}

func (fr *FeedbackRepository) GetFeedbackBySessionID(sessionID int) ([]*entity.Feedback, error) {
	var feedbacks []*entity.Feedback
	err := fr.db.Preload("User").Where("session_id = ?", sessionID).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (fr *FeedbackRepository) DeleteFeedbackInSession(feedbackID int) error {
	var feedback entity.Feedback

	err := fr.db.Debug().Model(&feedback).Where("feedback_id = ?", feedbackID).Delete(&feedback).Error
	if err != nil {
		return err
	}

	return nil
}
