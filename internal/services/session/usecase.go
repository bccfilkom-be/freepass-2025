package session

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"

	"github.com/gin-gonic/gin"
)

type SessionUsecase interface {
	GetAllSession(ctx *gin.Context) ([]dto.GetAllSessionResponse, error)
	UpdateSession(ctx *gin.Context, req *dto.UpdateSessionRequest) error
}
