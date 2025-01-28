package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/constant"
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
	proposalRouter.GET("/", middleware.RequireAuth, handler.GetAllProposal)
	proposalRouter.POST("/", middleware.RequireAuth, handler.CreateSessionProposal)

	proposalRouter.GET("/:sessionId", middleware.RequireAuth, handler.GetProposalDetail)
	proposalRouter.PATCH("/:sessionId", middleware.RequireAuth, handler.UpdateSessionProposal)
	proposalRouter.DELETE("/:sessionId", middleware.RequireAuth, handler.DeleteProposal)

	proposalRouter.PUT("/:sessionId/approve", middleware.RequireAuth,
		middleware.RequireRoles(constant.ROLE_COORDINATOR), handler.ApproveProposal)
	proposalRouter.PUT("/:sessionId/decline", middleware.RequireAuth,
		middleware.RequireRoles(constant.ROLE_COORDINATOR), handler.DeclineProposal)
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

func (v *ProposalDelivery) DeleteProposal(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		v.response.InternalServerError(ctx, "Invalid type of sessionId!")
		return
	}

	err = v.proposalUsecase.DeleteProposal(ctx, uint(sessionId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Proposal deleted!", 200)
}

func (v *ProposalDelivery) ApproveProposal(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		v.response.InternalServerError(ctx, "Invalid type of sessionId!")
		return
	}

	err = v.proposalUsecase.ApproveProposal(uint(sessionId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Proposal approved!", 200)
}

func (v *ProposalDelivery) DeclineProposal(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		v.response.InternalServerError(ctx, "Invalid type of sessionId!")
		return
	}

	var req *dto.DecliendProposalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	err = v.proposalUsecase.DeclineProposal(uint(sessionId), req)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Proposal declined!", 200)
}
