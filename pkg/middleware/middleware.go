package middleware

import (
	"freepass-bcc/internal/service"
	"freepass-bcc/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
	OnlyAdmin(ctx *gin.Context)
	OnlyEC(ctx *gin.Context)
	RegisteredOnly(ctx *gin.Context)
	Cors() gin.HandlerFunc
	Timeout() gin.HandlerFunc
}

type middleware struct {
	service *service.Service
	jwtAuth jwt.Interface
}

func Init(service *service.Service, jwtAuth jwt.Interface) Interface {
	return &middleware{
		service: service,
		jwtAuth: jwtAuth,
	}
}
