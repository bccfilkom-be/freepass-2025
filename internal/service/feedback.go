package service

import (
	"freepass-bcc/entity"
	"freepass-bcc/internal/repository"

	"github.com/google/uuid"
)

type IFeedbackService interface {
	AddFeedback(userID uuid.UUID, sessionID int, comment string) (*entity.Feedback, error)
	GetFeedbackBySessionID(sessionID int) ([]*entity.Feedback, error)
	DeleteFeedbackInSession(feedbackID int) error
}

type FeedbackService struct {
	FeedbackRepository repository.IFeedbackRepository
}

func NewFeedbackService(FeedbackRepository repository.IFeedbackRepository) IFeedbackService {
	return &FeedbackService{FeedbackRepository}
}

func (fs *FeedbackService) AddFeedback(userID uuid.UUID, sessionID int, comment string) (*entity.Feedback, error) {
	feedback := &entity.Feedback{
		UserID:    userID,
		SessionID: sessionID,
		Comment:   comment,
	}

	err := fs.FeedbackRepository.AddFeedback(feedback)
	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func (fs *FeedbackService) GetFeedbackBySessionID(sessionID int) ([]*entity.Feedback, error) {
	feedbacks, err := fs.FeedbackRepository.GetFeedbackBySessionID(sessionID)
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func (fs *FeedbackService) DeleteFeedbackInSession(feedbackID int) error {
	err := fs.FeedbackRepository.DeleteFeedbackInSession(feedbackID)
	if err != nil {
		return err
	}

	return nil
}
