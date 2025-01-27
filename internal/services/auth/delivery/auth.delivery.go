package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthDelivery struct {
	router      *gin.Engine
	response    response.ResponseHandler
	authUsecase auth.AuthUsecase
	validator   validator.ValidationService
}

func NewAuthDelivery(router *gin.Engine, authUsecase auth.AuthUsecase, response response.ResponseHandler, validator validator.ValidationService) {
	handler := AuthDelivery{
		router,
		response,
		authUsecase,
		validator,
	}

	authRouter := router.Group("/auth")
	authRouter.POST("/sign-up", handler.SignUp)
}

func (v *AuthDelivery) SignUp(ctx *gin.Context) {
	var req *dto.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	if err := v.authUsecase.SignUp(req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "User Created", http.StatusCreated)
}
