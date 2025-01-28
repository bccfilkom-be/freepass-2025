package registration

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
)

type RegistrationRepository interface {
	GetAllRegisteredSession(userId uint) ([]domain.SessionRegistration, error)
	Create(userId, sessionId uint) error
	GetBySessionId(sessionId uint) (domain.SessionRegistration, error)
	RegisteredSessionBeetweenDate(userId, sessionId uint) (domain.SessionRegistration, error)
}
