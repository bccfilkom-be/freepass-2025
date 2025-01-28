package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/proposal"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProposalDelivery struct {
	router          *gin.Engine
	proposalUsecase proposal.ProposalUsecase
	response        response.ResponseHandler
	validator       validator.ValidationService
}

func NewProposalDelivery(
	router *gin.Engine,
	proposalUsecase proposal.ProposalUsecase,
	response response.ResponseHandler,
	validator validator.ValidationService,
) {
	handler := &ProposalDelivery{
		router, proposalUsecase, response, validator,
	}

	proposalRouter := router.Group("/proposal")
	proposalRouter.POST("/", middleware.RequireAuth, handler.CreateSessionProposal)
	proposalRouter.GET("/", middleware.RequireAuth, handler.GetAllProposal)
	proposalRouter.PATCH("/:sessionId", middleware.RequireAuth, handler.UpdateSessionProposal)
	proposalRouter.GET("/:sessionId", middleware.RequireAuth, handler.GetProposalDetail)
}

func (v *ProposalDelivery) CreateSessionProposal(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	var req *dto.CreateProposalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	err := v.proposalUsecase.CreateProposal(userId.(uint), req)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Session Proposal created!", 200)
}

func (v *ProposalDelivery) UpdateSessionProposal(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	param := ctx.Param("sessionId")

	sessionId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		v.response.InternalServerError(ctx, "Invalid type of sessionId!")
		return
	}

	var req *dto.UpdateProposalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	err = v.proposalUsecase.UpdateProposal(uint(sessionId), userId.(uint), req)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Session Proposal Updated!", 200)
}

func (v *ProposalDelivery) GetAllProposal(ctx *gin.Context) {
	res, err := v.proposalUsecase.GetAllProposal(ctx)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Proposal found!", 200)
}

func (v *ProposalDelivery) GetProposalDetail(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		v.response.InternalServerError(ctx, "Invalid type of sessionId!")
		return
	}

	res, err := v.proposalUsecase.GetProposalDetail(ctx, uint(sessionId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Proposal found!", 200)
}
