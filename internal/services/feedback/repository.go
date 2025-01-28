package feedback

import "jevvonn/bcc-be-freepass-2025/internal/models/domain"

type FeedbackRepository interface {
	GetAllBySessionId(sessionId uint) ([]domain.SessionFeedback, error)
	Create(data domain.SessionFeedback) error
}
