package usecase

import (
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"
)

type SessionUsecase struct {
	sessionRepo session.SessionRepository
}

func NewSessionUsecase(sessionRepo session.SessionRepository) session.SessionUsecase {
	return &SessionUsecase{sessionRepo}
}

func (u *SessionUsecase) GetAllSession() ([]domain.Session, error) {
	// userId := ctx.Query("userId")
	// status := ctx.Query("status")

	// if userId != "" {
	// 	userId, err := strconv.ParseInt(ctx.Query("userId"), 10, 32)
	// 	if err == nil {
	// 		filter.UserID = uint(userId)
	// 	}
	// }

	// if status != "" && slices.Contains(constant.ROLE_ARRAY, status) {
	// 	filter.Status = status
	// }

	return u.sessionRepo.GetAll(session.SessionFilter{
		Status: constant.STATUS_SESSION_ACCEPTED,
	})
}
