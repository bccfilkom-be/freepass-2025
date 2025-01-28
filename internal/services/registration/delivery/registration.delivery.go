package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/services/registration"

	"github.com/gin-gonic/gin"
)

type RegistrationDelivery struct {
	router              *gin.Engine
	response            response.ResponseHandler
	registrationUsecase registration.RegistrationUsecase
	validator           validator.ValidationService
}

func NewRegistrationDelivery(
	router *gin.Engine,
	response response.ResponseHandler,
	registrationUsecase registration.RegistrationUsecase,
	validator validator.ValidationService,
) {
	handler := &RegistrationDelivery{
		router, response, registrationUsecase, validator,
	}

	sessionRouter := router.Group("/session")
	sessionRouter.GET("/registered", middleware.RequireAuth, handler.GetAllRegisteredSession)
	sessionRouter.POST("/:sessionId/register", middleware.RequireAuth, handler.RegisterSession)
}

func (v *RegistrationDelivery) GetAllRegisteredSession(ctx *gin.Context) {
	userId := ctx.GetUint("userId")

	sessions, err := v.registrationUsecase.GetAllRegisteredSession(userId)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, sessions, "Registered sessions", 200)
}

func (v *RegistrationDelivery) RegisterSession(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	param := ctx.Param("sessionId")

	sessionId, err := helper.StringToUint(param)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	err = v.registrationUsecase.RegisterSession(userId, sessionId)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "Session registered", 200)
}
