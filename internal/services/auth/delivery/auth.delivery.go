package delivery

import (
	"jevvonn/bcc-be-freepass-2025/internal/models/dto"
	"jevvonn/bcc-be-freepass-2025/internal/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthDelivery struct {
	router      *gin.Engine
	authUsecase auth.AuthUsecase
}

func NewAuthDelivery(router *gin.Engine, authUsecase auth.AuthUsecase) {
	v := AuthDelivery{
		router,
		authUsecase,
	}

	authRouter := router.Group("/auth")
	authRouter.POST("/sign-up", v.SignUp)
}

func (v *AuthDelivery) SignUp(ctx *gin.Context) {
	var req *dto.SignUpRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	err := v.authUsecase.SignUp(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User Created",
	})
}
