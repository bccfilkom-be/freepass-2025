package session

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"time"
)

type SessionFilter struct {
	UserID uint
	Status string
}

type SessionRepository interface {
	Create(data domain.Session) error
	GetAll(filter SessionFilter) ([]domain.Session, error)
	GetAllBetwenDate(startDate, endDate time.Time, filter SessionFilter) ([]domain.Session, error)
}
