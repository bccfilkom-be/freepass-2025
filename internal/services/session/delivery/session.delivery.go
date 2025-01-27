package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/session"

	"github.com/gin-gonic/gin"
)

type SessionDelivery struct {
	router         *gin.Engine
	sessionUsecase session.SessionUsecase
	response       response.ResponseHandler
	validator      validator.ValidationService
}

func NewSessionDelivery(
	router *gin.Engine,
	sessionUsecase session.SessionUsecase,
	response response.ResponseHandler,
	validator validator.ValidationService,
) {
	handler := &SessionDelivery{
		router, sessionUsecase, response, validator,
	}

	sessionRouter := router.Group("/session")
	sessionRouter.POST("/", middleware.RequireAuth, handler.CreateSession)
}

func (v *SessionDelivery) CreateSession(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	var req *dto.CreateSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	err := v.sessionUsecase.CreateSession(userId.(uint), req)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Session created!", 200)
}
