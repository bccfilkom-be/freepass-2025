package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"
	"log"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	router      *gin.Engine
	userUsecase user.UserUsecase
	response    response.ResponseHandler
	validator   validator.ValidationService
}

func NewUserDelivery(
	router *gin.Engine,
	userUsecase user.UserUsecase,
	response response.ResponseHandler,
	validator validator.ValidationService,
) {
	handler := &UserDelivery{
		router, userUsecase, response, validator,
	}

	// Profile Routing
	router.GET("/profile", middleware.RequireAuth, handler.GetUserProfile)
	router.PATCH("/profile", middleware.RequireAuth, handler.UpdateUserProfile)

	// User Routing
	userRouter := router.Group("/user")
	userRouter.GET("/:id", middleware.RequireAuth, handler.GetUserDetail)
}

func (v *UserDelivery) GetUserProfile(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")

	response, err := v.userUsecase.GetUserProfile(userId.(uint))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, response, "User profile found!", 200)
}

func (v *UserDelivery) GetUserDetail(ctx *gin.Context) {
	param := ctx.Param("id")
	userId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
		return
	}

	response, err := v.userUsecase.GetUserDetail(uint(userId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, response, "User detail found!", 200)
}

func (v *UserDelivery) UpdateUserProfile(ctx *gin.Context) {
	var req *dto.UpdateUserProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	userId, _ := ctx.Get("userId")
	err := v.userUsecase.UpdateUserProfile(userId.(uint), req)
	if err != nil {
		log.Fatal(err)

		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "User profile updated!", 200)
}
