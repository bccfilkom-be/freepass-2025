package middleware

import (
	jwt_helper "jevvonn/bcc-be-freepass-2025/internal/helper/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuth(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You're not login yet!",
		})
		return
	}

	token := strings.Split(bearerToken, " ")[1]

	claims, err := jwt_helper.ParseJWTToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Set("userId", uint(claims["sub"].(float64)))
	ctx.Set("role", claims["role"])
	ctx.Next()
}
