package proposal

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"

	"github.com/gin-gonic/gin"
)

type ProposalUsecase interface {
	CreateProposal(userId uint, req *dto.CreateProposalRequest) error
	GetAllProposal(ctx *gin.Context) ([]dto.GetAllProposalResponse, error)
	UpdateProposal(sessionId, userId uint, req *dto.UpdateProposalRequest) error
}
