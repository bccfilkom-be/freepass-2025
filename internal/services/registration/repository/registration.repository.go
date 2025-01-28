package repository

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/services/registration"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"

	"gorm.io/gorm"
)

type RegistrationRepository struct {
	db          *gorm.DB
	sessionRepo session.SessionRepository
}

func NewRegistrationRepository(db *gorm.DB, sessionRepo session.SessionRepository) registration.RegistrationRepository {
	return &RegistrationRepository{db, sessionRepo}
}

func (v *RegistrationRepository) GetAllRegisteredSession(userId uint) ([]domain.SessionRegistration, error) {
	var data []domain.SessionRegistration
	err := v.db.Preload("Session").Where("user_id = ?", userId).Find(&data).Error
	return data, err
}

func (v *RegistrationRepository) Create(userId, sessionId uint) error {
	return v.db.Create(&domain.SessionRegistration{
		UserID:    userId,
		SessionID: sessionId,
	}).Error
}

func (v *RegistrationRepository) GetBySessionId(sessionId uint) (domain.SessionRegistration, error) {
	var registration domain.SessionRegistration
	err := v.db.Where("session_id = ?", sessionId).First(&registration).Error
	return registration, err
}

func (v *RegistrationRepository) RegisteredSessionBeetweenDate(userId, sessionId uint) (domain.SessionRegistration, error) {
	session, err := v.sessionRepo.GetById(sessionId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.SessionRegistration{}, errors.New("Session not found!")
		} else {
			return domain.SessionRegistration{}, err
		}
	}

	var sessionsInBetween []domain.Session
	err = v.db.Model(&domain.Session{}).Where(
		"(? BETWEEN session_start_date AND session_end_date OR ? BETWEEN session_start_date AND session_end_date OR (session_start_date >= ? AND session_end_date <= ?))",
		session.SessionStartDate, session.SessionEndDate, session.SessionStartDate, session.SessionEndDate,
	).Not(session.ID).Where("status = ?", constant.STATUS_SESSION_ACCEPTED).Select("id").Find(&sessionsInBetween).Error

	if err != nil {
		return domain.SessionRegistration{}, err
	}

	IDs := []uint{}
	for _, s := range sessionsInBetween {
		IDs = append(IDs, s.ID)
	}

	var registration domain.SessionRegistration
	err = v.db.Preload("Session").Where("user_id = ? AND session_id IN ?", userId, IDs).First(&registration).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.SessionRegistration{}, nil
		} else {
			return domain.SessionRegistration{}, err
		}
	}
	return registration, err
}
