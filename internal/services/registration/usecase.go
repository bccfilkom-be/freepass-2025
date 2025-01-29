package registration

import "jevvonn/bcc-be-freepass-2025/internal/models/dto"

type RegistrationUsecase interface {
	RegisterSession(userId, sessionId uint) error
	GetAllRegisteredSession(userId uint) ([]dto.GetSessionRegistrationResponse, error)
}
