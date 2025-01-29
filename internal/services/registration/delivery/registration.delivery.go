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

	sessionRouter := router.Group("/api/session")
	sessionRouter.GET("/registered", middleware.RequireAuth, handler.GetAllRegisteredSession)
	sessionRouter.POST("/:sessionId/register", middleware.RequireAuth, handler.RegisterSession)
}

// @title 			Get All Registered Session
//
//	@Tags			Session Registration
//	@Summary		Get All Registered Session
//	@Description	Get All Registered Session
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.JSONResponseModel{data=[]dto.GetSessionRegistrationResponse}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/session/registered [get]
func (v *RegistrationDelivery) GetAllRegisteredSession(ctx *gin.Context) {
	userId := ctx.GetUint("userId")

	sessions, err := v.registrationUsecase.GetAllRegisteredSession(userId)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, sessions, "Registered sessions", 200)
}

// @title 			Register Into Session
//
//	@Tags			Session Registration
//	@Summary		Register Into Session
//	@Description	Register Into Session
//	@Accept			json
//	@Produce		json
//
// @Param 			id path int true "Session ID"
//
//	@Success		200		{object}	response.JSONResponseModel{data=nil}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//
// @Security BearerAuth
//
//	@Router			/api/session/{id}/register [post]
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
