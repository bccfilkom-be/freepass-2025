package session

import "jevvonn/bcc-be-freepass-2025/internal/models/domain"

type SessionRepository interface {
	Create(data domain.Session) error
}
