package usecase

import (
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
)

type SessionUsecase struct {
	sessionRepo session.SessionRepository
}

func NewSessionUsecase(sessionRepo session.SessionRepository) session.SessionUsecase {
	return &SessionUsecase{sessionRepo}
}
