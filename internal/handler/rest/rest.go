package rest

import (
	"fmt"
	"freepass-bcc/internal/service"
	"freepass-bcc/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Cors(), r.middleware.Timeout())
	routerGroup := r.router.Group("api/v1")
	routerGroup.POST("/register", r.Register)
	routerGroup.POST("/login", r.Login)

	user := routerGroup.Group("/user")
	user.Use(r.middleware.AuthenticateUser, r.middleware.RegisteredOnly)
	user.PATCH("/update-profile/:userID", r.UpdateProfile)
	user.GET("/session", r.GetAllSession)
	user.GET("/feedbacks", r.GetFeedbackBySessionID)
	user.GET("/user-profile", r.GetUserByID)
	user.POST("/add-feedback", r.AddFeedback)
	user.POST("/create-proposal", r.CreateProposal)
	user.POST("/update-proposal", r.UpdateProposal)
	user.POST("/register-session", r.RegisterUserToSession)
	user.DELETE("/delete-proposal", r.DeleteProposal)

	event := routerGroup.Group("/ec")
	event.Use(r.middleware.AuthenticateUser, r.middleware.OnlyEC)
	event.GET("/get-proposal", r.GetAllProposal)
	event.PUT("/proposal/:proposalID/approve", r.ApprovedProposal)
	event.PATCH("/proposal/:proposalID/reject", r.RejectedProsal)
	event.DELETE("/sessions/:sessionID", r.DeleteSession)
	event.DELETE("/delete-feedbacks/", r.DeleteFeedbackInSession)

	admin := routerGroup.Group("/admin")
	admin.Use(r.middleware.AuthenticateUser, r.middleware.OnlyAdmin)
	admin.PATCH("/add-ec/:userID", r.AddNewEC)
	admin.PATCH("/remove-role/:userID", r.RemoveRole)

}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
