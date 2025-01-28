package session

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"time"
)

type SessionFilter struct {
	UserID    uint
	Status    string
	IncludeID []uint
	ExcludeID []uint
}

type SessionRepository interface {
	Create(data domain.Session) error
	GetAll(filter SessionFilter) ([]domain.Session, error)
	GetAllBetwenDate(startDate, endDate time.Time, filter SessionFilter) ([]domain.Session, error)
	Update(data domain.Session) error
	GetById(id uint) (domain.Session, error)
	DateInBetweenSession(startDate, endDate time.Time, filter SessionFilter) error
}
