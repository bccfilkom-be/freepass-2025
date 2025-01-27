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
	ctx.JSON(200, "All Good!")
}
