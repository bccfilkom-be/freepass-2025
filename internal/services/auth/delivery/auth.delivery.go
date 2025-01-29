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

	authRouter := router.Group("/api/auth")
	authRouter.POST("/sign-up", handler.SignUp)
	authRouter.POST("/sign-in", handler.SignIn)
}

// @title 			Sign Up New User
//
//	@Tags			Auth
//	@Summary		Sign Up New User
//	@Description	Sign Up New User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.SignUpRequest	true	"Sign Up Request"
//	@Success		200		{object}	response.JSONResponseModel{data=nil}
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//	@Router			/api/auth/sign-up [post]
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

// @title 			Sign In User
//
//	@Tags			Auth
//	@Summary		Sign In User
//	@Description	Sign In User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.SignInRequest	true	"Sign In Request"
//	@Success		200		{object}	response.JSONResponseModel{data=dto.SignInResponse} "User Logged In and JWT Token Generated"
//	@Failure		400		{object}	response.JSONResponseModel{data=nil}
//	@Failure		500		{object}	response.JSONResponseModel{data=nil}
//	@Router			/api/auth/sign-in [post]
func (v *AuthDelivery) SignIn(ctx *gin.Context) {
	var req *dto.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	tokenData, err := v.authUsecase.SignIn(req)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, tokenData, "User Logged In", http.StatusCreated)
}
