package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/helper"
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	"jevvonn/bcc-be-freepass-2025/internal/middleware"
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/user"

	"github.com/gin-gonic/gin"
)

type AdminDelivery struct {
	router      *gin.Engine
	userUsecase user.AdminUsecase
	response    response.ResponseHandler
	validator   validator.ValidationService
}

func NewAdminDelivery(
	router *gin.Engine,
	userUsecase user.AdminUsecase,
	response response.ResponseHandler,
	validator validator.ValidationService,
) {
	handler := &AdminDelivery{
		router, userUsecase, response, validator,
	}

	userRouter := router.Group("/user")
	userRouter.DELETE("/:id", middleware.RequireAuth, middleware.RequireRoles(constant.ROLE_ADMIN), handler.DeleteUser)
	userRouter.PUT("/:id/update-role", middleware.RequireAuth, middleware.RequireRoles(constant.ROLE_ADMIN), handler.UpdateRole)
}

func (v *AdminDelivery) DeleteUser(ctx *gin.Context) {
	param := ctx.Param("id")
	userId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
		return
	}

	err = v.userUsecase.DeleteUser(userId)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "User deleted!", 200)
}

func (v *AdminDelivery) UpdateRole(ctx *gin.Context) {
	var req dto.UpdateUserRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	if errorsData, err := v.validator.Validate(req); err != nil {
		v.response.BadRequest(ctx, errorsData, err.Error())
		return
	}

	param := ctx.Param("id")
	userId, err := helper.StringToUint(param)
	if err != nil {
		v.response.InternalServerError(ctx, err.Error())
		return
	}

	err = v.userUsecase.UpdateRole(userId, &req)
	if err != nil {
		v.response.BadRequest(ctx, nil, err.Error())
		return
	}

	v.response.OK(ctx, nil, "User role updated!", 200)
}
