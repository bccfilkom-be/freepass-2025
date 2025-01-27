package api

import (
	"fmt"
	"jevvonn/bcc-be-freepass-2025/internal/config"
	"jevvonn/bcc-be-freepass-2025/internal/database"
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	auth_delivery "jevvonn/bcc-be-freepass-2025/internal/services/auth/delivery"
	"jevvonn/bcc-be-freepass-2025/internal/services/auth/usecase"
	user_repository "jevvonn/bcc-be-freepass-2025/internal/services/user/repository"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer() {
	config := config.GetConfig()
	router := gin.Default()

	host := config.GetString("host")
	port := config.GetString("port")

	db := database.NewDatabase()
	validator := validator.NewValidator()

	response := response.NewResponseHandler()

	// Repository
	userRepo := user_repository.NewUserRepository(db)

	// Usecase
	authUsecase := usecase.NewAuthUsecase(userRepo)

	// Delivery
	auth_delivery.NewAuthDelivery(router, authUsecase, response, validator)

	router.NoRoute(func(ctx *gin.Context) {
		response.NotFound(ctx)
	})
	router.Run(fmt.Sprintf("%s:%s", host, port))
}
