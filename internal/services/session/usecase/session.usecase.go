package usecase

import (
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
	"time"
)

type SessionUsecase struct {
	sessionRepo session.SessionRepository
}

func NewSessionUsecase(sessionRepo session.SessionRepository) session.SessionUsecase {
	return &SessionUsecase{sessionRepo}
}

func (v *SessionUsecase) GetAllSession() ([]dto.GetAllSessionResponse, error) {
	results, err := v.sessionRepo.GetAll(session.SessionFilter{
		Status: []string{constant.STATUS_SESSION_ACCEPTED},
	})

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
