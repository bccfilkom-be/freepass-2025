package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/registration"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
	"time"
)

type RegistrationUsecase struct {
	registrationRepo registration.RegistrationRepository
	sessionRepo      session.SessionRepository
}

func NewRegistrationUsecase(registrationRepo registration.RegistrationRepository, sessionRepo session.SessionRepository) registration.RegistrationUsecase {
	return &RegistrationUsecase{registrationRepo, sessionRepo}
}

func (v *RegistrationUsecase) RegisterSession(userId, sessionId uint) error {
	_, err := v.registrationRepo.GetBySessionId(sessionId)
	if err == nil {
		return errors.New("Session already registered!")
	}

	registeredSession, err := v.registrationRepo.RegisteredSessionBeetweenDate(userId, sessionId)
	if err != nil {
		return err
	}

	if registeredSession.Session.ID != 0 {
		return errors.New("Session already registered in that date range!")
	}

	return v.registrationRepo.Create(userId, sessionId)
}

func (v *RegistrationUsecase) GetAllRegisteredSession(userId uint) ([]dto.GetSessionRegistrationResponse, error) {
	registeredSessions, err := v.registrationRepo.GetAllRegisteredSession(userId)
	if err != nil {
		return nil, err
	}

	var sessions []dto.GetSessionRegistrationResponse
	for _, registeredSession := range registeredSessions {
		sessions = append(sessions, dto.GetSessionRegistrationResponse{
			ID: registeredSession.ID,
			Session: dto.GetAllSessionResponse{
				ID:                    registeredSession.Session.ID,
				Title:                 registeredSession.Session.Title,
				Description:           registeredSession.Session.Description,
				RegistrationStartDate: registeredSession.Session.RegistrationStartDate.Format(time.RFC3339),
				RegistrationEndDate:   registeredSession.Session.RegistrationEndDate.Format(time.RFC3339),

				SessionStartDate: registeredSession.Session.SessionStartDate.Format(time.RFC3339),
				SessionEndDate:   registeredSession.Session.SessionEndDate.Format(time.RFC3339),

				MaxSeat: registeredSession.Session.MaxSeat,

				CreatedAt: registeredSession.Session.CreatedAt.Format(time.RFC3339),
				UpdatedAt: registeredSession.Session.UpdatedAt.Format(time.RFC3339),
			},
		})
	}

	return sessions, nil
}
