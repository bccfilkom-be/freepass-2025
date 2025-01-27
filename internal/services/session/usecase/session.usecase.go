package usecase

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
	"time"
)

type SesionUsecase struct {
	sessionRepo session.SessionRepository
}

func NewSessionUsecase(sessionRepo session.SessionRepository) session.SessionUsecase {
	return &SesionUsecase{sessionRepo}
}

func (u *SesionUsecase) CreateSession(userId uint, req *dto.CreateSessionRequest) error {
	registrationStarDate, err := helper.StringISOToDateTime(req.RegistrationStartDate)
	if err != nil {
		return err
	}

	registrationEndDate, err := helper.StringISOToDateTime(req.RegistrationEndDate)
	if err != nil {
		return err
	}

	sessionStartDate, err := helper.StringISOToDateTime(req.SessionStartDate)
	if err != nil {
		return err
	}

	sessionEndDate, err := helper.StringISOToDateTime(req.SessionEndDate)
	if err != nil {
		return err
	}

	if registrationStarDate.Before(time.Now()) {
		return errors.New("Registration start date should be after today!")
	}

	if registrationStarDate.After(registrationEndDate) {
		return errors.New("Registration start date should be before the registration end date!")
	}

	if sessionStartDate.Before(registrationEndDate) {
		return errors.New("Session start date should be after the registration end date!")
	}

	if sessionStartDate.After(sessionEndDate) {
		return errors.New("Session start date should be before the session start date!")
	}

	data := domain.Session{
		UserID:                userId,
		Title:                 req.Title,
		Description:           req.Description,
		RegistrationStartDate: registrationStarDate,
		RegistrationEndDate:   registrationEndDate,

		SessionStartDate: sessionStartDate,
		SessionEndDate:   sessionEndDate,
		MaxSeat:          req.MaxSeat,
	}

	return u.sessionRepo.Create(data)
}
