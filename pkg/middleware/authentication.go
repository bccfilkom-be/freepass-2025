package middleware

import (
	"errors"
	"freepass-bcc/model"
	"freepass-bcc/pkg/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthenticateUser(ctx *gin.Context) {
	bearer := ctx.GetHeader("Authorization")
	if bearer == "" {
		response.Error(ctx, http.StatusUnauthorized, "empty token", errors.New(""))
		return
	}

	token := strings.Split(bearer, " ")[1]
	userID, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to validate token", err)
		ctx.Abort()
		return
	}

	user, err := m.service.UserService.GetUser(model.UserParam{
		UserID: userID,
	})
	if err != nil {
		response.Error(ctx, http.StatusUnauthorized, "failed to get user", err)
		ctx.Abort()
		return
	}

	ctx.Set("user", user)
	ctx.Next()
}
