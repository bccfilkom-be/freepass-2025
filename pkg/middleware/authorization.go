package middleware

import (
	"errors"
	"freepass-bcc/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) OnlyAdmin(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed get login user", err)
		ctx.Abort()
		return
	}

	if user.RoleID != 1 {
		response.Error(ctx, http.StatusForbidden, "this endpoint cannot be access", errors.New("user dont have access"))
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (m *middleware) OnlyEC(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed get login user", err)
		ctx.Abort()
		return
	}

	if user.RoleID != 2 {
		response.Error(ctx, http.StatusForbidden, "this endpoint cannot be access", errors.New("user dont have access"))
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (m *middleware) RegisteredOnly(ctx *gin.Context) {
	user, err := m.jwtAuth.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, http.StatusForbidden, "failed get login user", err)
		ctx.Abort()
		return
	}

	if user.RoleID == 4 {
		response.Error(ctx, http.StatusForbidden, "this endpoint cannot be access", errors.New("user dont have access"))
		ctx.Abort()
		return
	}
	ctx.Next()
}
