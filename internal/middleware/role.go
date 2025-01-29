package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func RequireRoles(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exist := ctx.Get("role")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You're not login yet!",
			})
			return
		}

		if !slices.Contains(roles, userRole.(string)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "You're not authorized to access this resource!",
			})
			return
		}

		ctx.Next()
	}
}
