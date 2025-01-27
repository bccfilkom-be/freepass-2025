package session

import "jevvonn/bcc-be-freepass-2025/internal/models/dto"

type SessionUsecase interface {
	GetAllSession() ([]dto.GetAllSessionResponse, error)
}
