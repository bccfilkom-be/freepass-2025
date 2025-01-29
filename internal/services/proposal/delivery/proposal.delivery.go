package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/proposal"

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

	proposalRouter := router.Group("/api/proposal")
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

// @title 			Create New Session Proposal
//
//	@Tags			Session Proposal
//	@Summary		Create New Session Proposal
//	@Description	**Create New Session Proposal**
//
// @Description
// @Description All date-time fields must follow the ISO 8601 format:
// @Description
// @Description **Date Format (ISO 8601):** (YYYY-MM-DD)T(HH:MM:SS)Z
// @Description Example: 2025-02-15T12:00:00Z
// @Description
//
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateProposalRequest	true	"Create Session Request"
//	@Success		200		{object}	response.JSONResponseModel{data=nil}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/proposal [post]
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

// @title 			Update Session Proposal
//
//	@Tags			Session Proposal
//	@Summary		Update Session Proposal
//	@Description	Update Session Proposal
//
// @Description
// @Description All date-time fields must follow the ISO 8601 format:
// @Description
// @Description **Date Format (ISO 8601):** (YYYY-MM-DD)T(HH:MM:SS)Z
// @Description Example: 2025-02-15T12:00:00Z
// @Description
//
//	@Accept			json
//	@Produce		json
//	@Param			id			path		int	true	"Session ID"
//	@Param			request	body		dto.UpdateProposalRequest	true	"Update Session Request"
//	@Success		200		{object}	response.JSONResponseModel{data=nil}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/proposal/{id} [patch]
func (v *ProposalDelivery) UpdateSessionProposal(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
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

// @title 			Get All Session Proposal
//
//	@Tags			Session Proposal
//	@Summary		Get All Session Proposal
//	@Description	Get All Session Proposal
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.JSONResponseModel{data=[]dto.GetProposalResponse}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/proposal [get]
func (v *ProposalDelivery) GetAllProposal(ctx *gin.Context) {
	res, err := v.proposalUsecase.GetAllProposal(ctx)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Proposal found!", 200)
}

// @title 			Get Detail Session Proposal
//
//	@Tags			Session Proposal
//	@Summary		Get Detail Session Proposal
//	@Description	Get Detail Session Proposal
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Session ID"
//	@Success		200		{object}	response.JSONResponseModel{data=dto.GetProposalResponse}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/proposal/{id} [get]
func (v *ProposalDelivery) GetProposalDetail(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
		return
	}

	res, err := v.proposalUsecase.GetProposalDetail(ctx, uint(sessionId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Proposal found!", 200)
}

// @title 			Delete Session Proposal
//
//	@Tags			Session Proposal
//	@Summary		Delete Session Proposal
//	@Description	Delete Session Proposal. Only It's own User session proposal can delete.
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Session ID"
//	@Success		200		{object}	response.JSONResponseModel{data=dto.GetProposalResponse}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/proposal/{id} [delete]
func (v *ProposalDelivery) DeleteProposal(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
		return
	}

	err = v.proposalUsecase.DeleteProposal(ctx, uint(sessionId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Proposal deleted!", 200)
}

// @title 			Approve Session Proposal
//
//	@Tags			Session Proposal
//	@Summary		Approve Session Proposal
//	@Description	Approve Session Proposal. Only Event Coordinator can approve.
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Session ID"
//	@Success		200		{object}	response.JSONResponseModel{data=nil}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/proposal/{id}/approve [put]
func (v *ProposalDelivery) ApproveProposal(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
		return
	}

	err = v.proposalUsecase.ApproveProposal(uint(sessionId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Proposal approved!", 200)
}

// @title 			Decline Session Proposal
//
//		@Tags			Session Proposal
//		@Summary		Decline Session Proposal
//		@Description	Decline Session Proposal. Only Event Coordinator can decline.
//		@Accept			json
//		@Produce		json
//		@Param			id		path		int	true	"Session ID"
//	 @Param			request	body		dto.DecliendProposalRequest	true	"Decline Proposal Request"
//		@Success		200		{object}	response.JSONResponseModel{data=nil}
//		@Failure		400		{object}	response.JSONResponseModel{data=nil}
//		@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/proposal/{id}/decline [put]
func (v *ProposalDelivery) DeclineProposal(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
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
