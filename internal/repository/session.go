package repository

import (
	"freepass-bcc/entity"

	"gorm.io/gorm"
)

type ISessionRepository interface {
	GetAllSession(limit, offset int) (*[]entity.Session, error)
	DeleteSession(sessionID int) error
}

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) ISessionRepository {
	return &SessionRepository{db}
}

func (sr *SessionRepository) GetAllSession(limit, offset int) (*[]entity.Session, error) {
	var session *[]entity.Session

	err := sr.db.Debug().Limit(limit).Offset(offset).Find(&session).Error
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (sr *SessionRepository) DeleteSession(sessionID int) error {
	var (
		feedback entity.Feedback
		session  entity.Session
	)

	err := sr.db.Debug().Where("session_id = ?", sessionID).Delete(&feedback).Error
	if err != nil {
		return err
	}

	err = sr.db.Debug().Where("session_id = ?", sessionID).Delete(&session).Error
	if err != nil {
		return err
	}

	return nil
}
