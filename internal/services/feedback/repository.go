package feedback

import "jevvonn/bcc-be-freepass-2025/internal/models/domain"

type FeedbackRepository interface {
	Create(data domain.SessionFeedback) error
}
