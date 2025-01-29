package repository

import (
	"freepass-bcc/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IRegisterRepository interface {
	CreateRegistration(register *entity.SessionRegistration) error
	GetSessionByID(sessionID int) (*entity.Session, error)
	GetUserRegistToday(userID uuid.UUID) (*entity.SessionRegistration, error)
}

type RegisterRepository struct {
	db *gorm.DB
}

func NewRegisterRepository(db *gorm.DB) IRegisterRepository {
	return &RegisterRepository{db}
}

func (rr *RegisterRepository) CreateRegistration(register *entity.SessionRegistration) error {
	err := rr.db.Debug().Create(register).Error
	if err != nil {
		return err
	}

	return nil
}

func (rr *RegisterRepository) GetSessionByID(sessionID int) (*entity.Session, error) {
	var session entity.Session
	err := rr.db.Debug().Where("session_id = ?", sessionID).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &session, nil
}

func (rr *RegisterRepository) GetUserRegistToday(userID uuid.UUID) (*entity.SessionRegistration, error) {
	var regist entity.SessionRegistration
	dayStart := time.Now().Truncate(24 * time.Hour)
	dayEnd := dayStart.Add(24 * time.Hour)

	err := rr.db.Debug().Where("user_id = ? AND timestamp BETWEEN ? AND ?", userID, dayStart, dayEnd).First(&regist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &regist, nil
}
