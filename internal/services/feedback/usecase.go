package feedback

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"

	"github.com/gin-gonic/gin"
)

type FeedbackUsecase interface {
	CreateFeedback(ctx *gin.Context, req *dto.CreateFeedbackRequest) error
}
