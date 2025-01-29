package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper"
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

	sessionRouter := router.Group("/api/session")
	sessionRouter.GET("/", middleware.RequireAuth, handler.GetAllSession)
	sessionRouter.PATCH("/:sessionId", middleware.RequireAuth, handler.UpdateSession)
	sessionRouter.DELETE("/:sessionId", middleware.RequireAuth, handler.DeleteSession)
}

func (v *SessionDelivery) GetAllSession(ctx *gin.Context) {
	res, err := v.sessionUsecase.GetAllSession(ctx)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Sessions found!", 200)
}

func (v *SessionDelivery) UpdateSession(ctx *gin.Context) {
	var req *dto.UpdateSessionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	if err := v.sessionUsecase.UpdateSession(ctx, req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Session updated!", 200)
}

func (v *SessionDelivery) DeleteSession(ctx *gin.Context) {
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
		return
	}

	err = v.sessionUsecase.DeleteSession(ctx, uint(sessionId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Session deleted!", 200)
}
