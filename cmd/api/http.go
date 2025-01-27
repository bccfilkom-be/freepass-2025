package api

import (
	"fmt"
	"jevvonn/bcc-be-freepass-2025/internal/config"
	"jevvonn/bcc-be-freepass-2025/internal/database"
	"jevvonn/bcc-be-freepass-2025/internal/helper/response"
	"jevvonn/bcc-be-freepass-2025/internal/helper/validator"
	auth_delivery "jevvonn/bcc-be-freepass-2025/internal/services/auth/delivery"
	proposal_delivery "jevvonn/bcc-be-freepass-2025/internal/services/proposal/delivery"
	session_delivery "jevvonn/bcc-be-freepass-2025/internal/services/session/delivery"
	user_delivery "jevvonn/bcc-be-freepass-2025/internal/services/user/delivery"

	auth_usecase "jevvonn/bcc-be-freepass-2025/internal/services/auth/usecase"
	proposal_usecase "jevvonn/bcc-be-freepass-2025/internal/services/proposal/usecase"
	session_usecase "jevvonn/bcc-be-freepass-2025/internal/services/session/usecase"
	user_usecase "jevvonn/bcc-be-freepass-2025/internal/services/user/usecase"

	session_repository "jevvonn/bcc-be-freepass-2025/internal/services/session/repository"
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
	sessionRepo := session_repository.NewSessionRepository(db)

	// Usecase
	authUsecase := auth_usecase.NewAuthUsecase(userRepo)
	userUsecase := user_usecase.NewUserUsecase(userRepo)
	sessionUsecase := session_usecase.NewSessionUsecase(sessionRepo)
	proposalUsecase := proposal_usecase.NewProposalUsecase(sessionRepo)

	// Delivery
	auth_delivery.NewAuthDelivery(router, authUsecase, response, validator)
	user_delivery.NewUserDelivery(router, userUsecase, response, validator)
	session_delivery.NewSessionDelivery(router, sessionUsecase, response, validator)
	proposal_delivery.NewProposalDelivery(router, proposalUsecase, response, validator)

	router.NoRoute(func(ctx *gin.Context) {
		response.NotFound(ctx)
	})
	router.Run(fmt.Sprintf("%s:%s", host, port))
}
