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

	sessionRoute := handler.router.Group("/session")
	sessionRoute.GET("/:sessionId/feedback", middleware.RequireAuth, handler.GetAllSessionFeedback)
	sessionRoute.POST("/:sessionId/feedback", middleware.RequireAuth, handler.CreateFeedback)
}

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

func (v *FeedbackDelivery) GetAllSessionFeedback(ctx *gin.Context) {
	res, err := v.feedbackUsecase.GetAllSessionFeedback(ctx)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Session feedback found!", 200)
}
