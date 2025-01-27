package session

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
)

type SessionUsecase interface {
	GetAllSession() ([]domain.Session, error)
}
