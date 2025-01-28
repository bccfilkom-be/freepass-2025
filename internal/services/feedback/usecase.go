package feedback

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"

	"github.com/gin-gonic/gin"
)

type FeedbackUsecase interface {
	GetAllSessionFeedback(ctx *gin.Context) ([]dto.GetFeedbackResponse, error)
	CreateFeedback(ctx *gin.Context, req *dto.CreateFeedbackRequest) error
	DeleteFeedback(ctx *gin.Context) error
}
