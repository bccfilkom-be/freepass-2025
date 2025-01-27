package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"
	"strconv"

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
	profileRouter := router.Group("/profile")
	profileRouter.GET("/", middleware.RequireAuth, handler.GetUserProfile)

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
	userId, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		v.response.InternalServerError(ctx, "Invalid type of userId!")
		return
	}

	response, err := v.userUsecase.GetUserDetail(uint(userId))
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, response, "User detail found!", 200)
}
