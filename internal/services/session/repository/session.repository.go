package repository

import (
	"errors"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
	"time"

	"gorm.io/gorm"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) session.SessionRepository {
	return &SessionRepository{db}
}

func (v *SessionRepository) Create(data domain.Session) error {
	return v.db.Create(&data).Error
}

func (v *SessionRepository) Update(data domain.Session) error {
	if data.ID == 0 {
		return errors.New("Session not found!")
	}

	return v.db.Updates(&data).Error
}

func (v *SessionRepository) GetAll(filter session.SessionFilter) ([]domain.Session, error) {
	var data []domain.Session
	query := v.db.Model(&domain.Session{})
	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if len(filter.Status) > 0 {
		query = query.Where("status IN ?", filter.Status)
	}

	err := query.Preload("User").Find(&data).Error
	return data, err
}

func (v *SessionRepository) GetById(id uint) (domain.Session, error) {
	var data domain.Session
	err := v.db.Preload("User").First(&data, id).Error
	return data, err
}

func (v *SessionRepository) DateInBetweenSession(startDate, endDate time.Time, filter session.SessionFilter) error {
	var data []domain.Session
	query := v.db.Model(&domain.Session{})
	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if len(filter.Status) > 0 {
		query = query.Where("status IN ?", filter.Status)
	}

	// Group the OR conditions properly
	query = query.Where(
		"(? BETWEEN session_start_date AND session_end_date OR ? BETWEEN session_start_date AND session_end_date OR (session_start_date >= ? AND session_end_date <= ?))",
		startDate, endDate, startDate, endDate,
	)

	if len(filter.ExcludeID) > 0 {
		query = query.Not(filter.ExcludeID)
	}

	err := query.Find(&data).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}

	if len(data) > 0 {
		return errors.New("Session already exist in that date range!")
	}

	return nil
}

func (v *SessionRepository) GetAllBetwenDate(startDate, endDate time.Time, filter session.SessionFilter) ([]domain.Session, error) {
	var data []domain.Session
	query := v.db.Model(&domain.Session{})
	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if len(filter.Status) > 0 {
		query = query.Where("status IN ?", filter.Status)
	}

	query = query.Where("session_start_date >= ? AND session_end_date <= ?", startDate, endDate)

	err := query.Preload("User").Find(&data).Error
	return data, err
}

func (v *SessionRepository) Delete(id uint) error {
	return v.db.Delete(&domain.Session{}, id).Error
}
