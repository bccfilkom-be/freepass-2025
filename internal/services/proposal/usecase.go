package proposal

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"

	"github.com/gin-gonic/gin"
)

type ProposalUsecase interface {
	CreateProposal(userId uint, req *dto.CreateProposalRequest) error
	GetAllProposal(ctx *gin.Context) ([]dto.GetProposalResponse, error)
	UpdateProposal(sessionId, userId uint, req *dto.UpdateProposalRequest) error
	GetProposalDetail(ctx *gin.Context, sessionId uint) (dto.GetProposalResponse, error)
	DeleteProposal(ctx *gin.Context, sessionId uint) error
	ApproveProposal(sessionId uint) error
	DeclineProposal(sessionId uint, req *dto.DecliendProposalRequest) error
}
