package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
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
	sessionRouter.GET("/", handler.GetAllSession)
}

func (v *SessionDelivery) GetAllSession(ctx *gin.Context) {
	res, err := v.sessionUsecase.GetAllSession(ctx)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, res, "Sessions found!", 200)
}
