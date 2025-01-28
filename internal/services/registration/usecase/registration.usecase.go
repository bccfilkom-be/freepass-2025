package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/services/registration"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
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
