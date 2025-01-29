package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionUsecase struct {
	sessionRepo session.SessionRepository
}

func NewSessionUsecase(sessionRepo session.SessionRepository) session.SessionUsecase {
	return &SessionUsecase{sessionRepo}
}

func (v *SessionUsecase) GetAllSession(ctx *gin.Context) ([]dto.GetAllSessionResponse, error) {
	filter := session.SessionFilter{
		Status: []string{constant.STATUS_SESSION_ACCEPTED},
	}

	userIdQuery := ctx.Query("userId")
	if userIdQuery != "" {
		userId, err := helper.StringToUint(userIdQuery)
		if err != nil {
			return []dto.GetAllSessionResponse{}, err
		}

		filter.UserID = userId
	}

	results, err := v.sessionRepo.GetAll(filter)

	if err != nil {
		return []dto.GetAllSessionResponse{}, err
	}

	var sessions []dto.GetAllSessionResponse
	for _, session := range results {
		sessions = append(sessions, dto.GetAllSessionResponse{
			ID:                    session.ID,
			Title:                 session.Title,
			Description:           session.Description,
			RegistrationStartDate: session.RegistrationStartDate.Format(time.RFC3339),
			RegistrationEndDate:   session.RegistrationEndDate.Format(time.RFC3339),

			SessionStartDate: session.SessionStartDate.Format(time.RFC3339),
			SessionEndDate:   session.SessionEndDate.Format(time.RFC3339),

			MaxSeat: session.MaxSeat,

			User: dto.GetUserDetailResponse{
				ID:    session.User.ID,
				Name:  session.User.Name,
				Email: session.User.Email,
			},

			CreatedAt: session.CreatedAt.Format(time.RFC3339),
			UpdatedAt: session.UpdatedAt.Format(time.RFC3339),
		})
	}

	return sessions, nil
}

func (v *SessionUsecase) UpdateSession(ctx *gin.Context, req *dto.UpdateSessionRequest) error {
	param := ctx.Param("sessionId")
	userId := ctx.GetUint("userId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		return err
	}

	sessionExist, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Session not found!")
		} else {
			return err
		}
	}

	if sessionExist.Status != constant.STATUS_SESSION_ACCEPTED {
		return errors.New("Session not found! Either the session is not accepted or already deleted.")
	}

	if sessionExist.UserID != userId {
		return errors.New("You are not authorized to update this session!")
	}

	dates, err := helper.ParseDatesFromRequest(
		req.RegistrationStartDate,
		req.RegistrationEndDate,
		req.SessionStartDate,
		req.SessionEndDate,
	)

	if err != nil {
		return err
	}

	if err := helper.ValidateDates(dates); err != nil {
		return err
	}

	if err := v.sessionRepo.DateInBetweenSession(dates.SessionStart, dates.SessionEnd, session.SessionFilter{
		UserID:    userId,
		ExcludeID: []uint{sessionId},
	}); err != nil {
		return err
	}

	data := domain.Session{
		ID:                    sessionId,
		Title:                 req.Title,
		Description:           req.Description,
		RegistrationStartDate: dates.RegistrationStart,
		RegistrationEndDate:   dates.RegistrationEnd,

		SessionStartDate: dates.SessionStart,
		SessionEndDate:   dates.SessionEnd,
		MaxSeat:          req.MaxSeat,
	}

	return v.sessionRepo.Update(data)
}

func (v *SessionUsecase) DeleteSession(ctx *gin.Context, sessionId uint) error {
	userId := ctx.GetUint("userId")
	role := ctx.GetString("role")

	session, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		return err
	}

	if session.UserID != userId && role != constant.ROLE_COORDINATOR {
		return errors.New("You are not authorized to delete this session!")
	}

	if session.Status != constant.STATUS_SESSION_ACCEPTED {
		return errors.New("Session not found! Either the session is not accepted or deleted.")
	}

	return v.sessionRepo.Delete(sessionId)
}
