package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/feedback"
	"jevvonn/bcc-be-freepass-2025/internal/services/registration"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FeedbackUsecase struct {
	registrationRepo registration.RegistrationRepository
	sessionRepo      session.SessionRepository
	feedbackRepo     feedback.FeedbackRepository
}

func NewFeedbackUsecase(registrationRepo registration.RegistrationRepository, sessionRepo session.SessionRepository, feedbackRepo feedback.FeedbackRepository) feedback.FeedbackUsecase {
	return &FeedbackUsecase{registrationRepo, sessionRepo, feedbackRepo}
}

func (u *FeedbackUsecase) GetAllSessionFeedback(ctx *gin.Context) ([]dto.GetFeedbackResponse, error) {
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		return nil, err
	}

	_, err = u.sessionRepo.GetById(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Session not found! Either session is not exist or already deleted")
		} else {
			return nil, err
		}
	}

	feedbacks, err := u.feedbackRepo.GetAllBySessionId(sessionId)
	if err != nil {
		return nil, err
	}

	var responses []dto.GetFeedbackResponse
	for _, feedback := range feedbacks {
		responses = append(responses, dto.GetFeedbackResponse{
			ID:        feedback.ID,
			Content:   feedback.Content,
			Rating:    feedback.Rating,
			CreatedAt: feedback.CreatedAt.Format(time.RFC3339),
			UpdatedAt: feedback.UpdatedAt.Format(time.RFC3339),
			User: dto.GetUserDetailResponse{
				ID:    feedback.User.ID,
				Name:  feedback.User.Name,
				Email: feedback.User.Email,
				Bio:   feedback.User.Bio,
			},
		})
	}

	return responses, nil
}

func (u *FeedbackUsecase) CreateFeedback(ctx *gin.Context, req *dto.CreateFeedbackRequest) error {
	userId := ctx.GetUint("userId")
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		return err
	}

	// Check if user already registered to the session
	_, err = u.registrationRepo.GetBySessionId(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("You're not registered to the session")
		} else {
			return err
		}
	}

	session, err := u.sessionRepo.GetById(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Session not found! Either session is not exist or already deleted")
		} else {
			return err
		}
	}

	if time.Now().Before(session.SessionStartDate) {
		return errors.New("Session hasn't started yet")
	}

	if time.Now().Before(session.SessionEndDate) {
		return errors.New("Wait until the session end to give feedback")
	}

	data := domain.SessionFeedback{
		UserID:    userId,
		SessionID: sessionId,
		Content:   req.Content,
		Rating:    req.Rating,
	}
	return u.feedbackRepo.Create(data)
}
