package registration

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
)

type RegistrationRepository interface {
	Create(userId, sessionId uint) error
	GetBySessionId(sessionId uint) (domain.SessionRegistration, error)
	RegisteredSessionBeetweenDate(userId, sessionId uint) (domain.SessionRegistration, error)
}
