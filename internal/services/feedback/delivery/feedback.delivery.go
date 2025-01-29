package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/feedback"

	"github.com/gin-gonic/gin"
)

type FeedbackDelivery struct {
	router          *gin.Engine
	response        response.ResponseHandler
	feedbackUsecase feedback.FeedbackUsecase
	validator       validator.ValidationService
}

func NewFeedbackDelivery(
	router *gin.Engine,
	response response.ResponseHandler,
	feedbackUsecase feedback.FeedbackUsecase,
	validator validator.ValidationService,
) {
	handler := &FeedbackDelivery{
		router, response, feedbackUsecase, validator,
	}

	sessionRoute := handler.router.Group("/api/session")
	sessionRoute.GET("/:sessionId/feedback", middleware.RequireAuth, handler.GetAllSessionFeedback)
	sessionRoute.POST("/:sessionId/feedback", middleware.RequireAuth, handler.CreateFeedback)
	sessionRoute.DELETE("/:sessionId/feedback/:feedbackId", middleware.RequireAuth, handler.DeleteFeedback)
}

// @title 			Create Session Feedback
//
//	@Tags			Session Feedback
//	@Summary		Create Session Feedback
//	@Description	Create Session Feedback
//	@Accept			json
//	@Produce		json
//
// @Param id path int true "Session ID"
// @Param request body dto.CreateFeedbackRequest true "Create Feedback Request"
//
//	@Success		200		{object}	response.JSONResponseModel{data=nil}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/session/{id}/feedback [post]
func (v *FeedbackDelivery) CreateFeedback(ctx *gin.Context) {
	var req dto.CreateFeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	err := v.feedbackUsecase.CreateFeedback(ctx, &req)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Feedback Created!", 201)
}

// @title 			Get Session Feedback
//
//	@Tags			Session Feedback
//	@Summary		Get Session Feedback
//	@Description	Get Session Feedback
//	@Accept			json
//	@Produce		json
//
// @Param id path int true "Session ID"
//
//	@Success		200		{object}	response.JSONResponseModel{data=[]dto.GetFeedbackResponse}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/session/{id}/feedback [get]
func (v *FeedbackDelivery) GetAllSessionFeedback(ctx *gin.Context) {
	res, err := v.feedbackUsecase.GetAllSessionFeedback(ctx)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Session feedback found!", 200)
}

// @title 			Delete Feedback
//
//	@Tags			Session Feedback
//	@Summary		Delete Feedback
//	@Description	Delete Feedback.  Only It's own User or Event Coordinator session feedback can delete
//	@Accept			json
//	@Produce		json
//
// @Param sessionId path int true "Session ID"
// @Param feedbackId path int true "Feedback ID"
//
//	@Success		200		{object}	response.JSONResponseModel{data=nil}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/session/{sessionId}/feedback/{feedbackId} [delete]
func (v *FeedbackDelivery) DeleteFeedback(ctx *gin.Context) {
	err := v.feedbackUsecase.DeleteFeedback(ctx)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Feedback deleted!", 200)
}
