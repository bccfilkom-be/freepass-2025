package service

import (
	"errors"
	"freepass-bcc/entity"
	"freepass-bcc/internal/repository"
	"time"

	"github.com/google/uuid"
)

type ISessionService interface {
	GetAllSession(page int) (*[]entity.Session, error)
	DeleteSession(sessionID int) error
	RegisterUserToSession(userID uuid.UUID, sessionID int) (*entity.SessionRegistration, error)
}

type SessionService struct {
	SessionRepository  repository.ISessionRepository
	RegisterRepository repository.IRegisterRepository
}

func NewSessionService(sessionRepository repository.ISessionRepository, registerRepository repository.IRegisterRepository) ISessionService {
	return &SessionService{
		SessionRepository:  sessionRepository,
		RegisterRepository: registerRepository,
	}
}

func (ss *SessionService) GetAllSession(page int) (*[]entity.Session, error) {
	limit := 4
	offset := (page - 1) * limit

	session, err := ss.SessionRepository.GetAllSession(limit, offset)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (ss *SessionService) DeleteSession(sessionID int) error {
	err := ss.SessionRepository.DeleteSession(sessionID)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SessionService) RegisterUserToSession(userID uuid.UUID, sessionID int) (*entity.SessionRegistration, error) {
	existRegist, err := ss.RegisterRepository.GetUserRegistToday(userID)
	if err != nil {
		return nil, err
	}

	if existRegist != nil {
		return nil, errors.New("user already registered")
	}

	session, err := ss.RegisterRepository.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	if session.AvailableSeats <= 0 {
		return nil, errors.New("session full")
	}
	session.AvailableSeats -= 1

	registration := &entity.SessionRegistration{
		UserID:    userID,
		SessionID: sessionID,
		Timestamp: time.Now(),
	}

	err = ss.RegisterRepository.CreateRegistration(registration)
	if err != nil {
		return nil, err
	}

	return registration, nil
}
