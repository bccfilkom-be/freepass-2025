package main

import (
	"context"
	"log"
	"os"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/litegral/freepass-2025/internal/controller"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/middleware"
	"github.com/litegral/freepass-2025/internal/service"
)

func main() {
	// Initialize environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v\n", err)
	}

	// Initialize database connection
	dbConn, err := pgx.Connect(context.Background(), os.Getenv("APP_DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbConn.Close(context.Background())

	// Initialize SQLC queries
	queries := db.New(dbConn)

	// Initialize services
	emailService, err := service.NewEmailService()
	if err != nil {
		log.Fatalf("Unable to initialize email service: %v\n", err)
	}
	userService := service.NewUserService(queries, emailService)
	sessionService := service.NewSessionService(queries)
	coordinatorService := service.NewCoordinatorService(queries)

	// Initialize controllers
	userController := controller.NewUserController(userService)
	sessionController := controller.NewSessionController(sessionService)
	coordinatorController := controller.NewCoordinatorController(coordinatorService)

	// Initialize router with custom error handler
	s := fuego.NewServer(
		fuego.WithAddr(os.Getenv("APP_URL")),
	)

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (fuego.HTML, error) {
		return "Conflux API is running! Check out <a href='http://localhost:8080/swagger/index.html'>the API Docs!</a>", nil
	})

	v1 := fuego.Group(s, "/v1")

	fuego.Post(v1, "/users", userController.CreateUser, option.Tags("Auth"))
	fuego.Get(v1, "/verify-email",
		userController.VerifyEmail,
		option.Query("token", "Token for email verification", param.Required()),
		option.Query("email", "Email for email verification", param.Required()),
		option.Tags("Auth"),
	)
	fuego.Post(v1, "/login", userController.Login, option.Tags("Auth"))

	// Protected routes group
	protected := fuego.Group(v1, "")
	fuego.Use(protected, middleware.AuthMiddleware(queries))

	fuego.Put(protected, "/profile", userController.UpdateProfile,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Update user profile"),
		option.Summary("Update Profile"),
		option.Tags("User"),
	)

	fuego.Get(protected, "/profile", userController.GetProfile,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get current user's profile"),
		option.Summary("Get Profile"),
		option.Tags("User"),
	)

	fuego.Get(protected, "/users/{id}", userController.GetUserProfile,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get user profile by ID"),
		option.Summary("Get User Profile"),
		option.Tags("User"),
		option.Path("id", "User ID", param.Required()),
	)

	// Session proposal routes
	fuego.Post(protected, "/sessions", sessionController.CreateProposal,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Create a new session proposal"),
		option.Summary("Create Session Proposal"),
		option.Tags("Session Proposal"),
	)

	fuego.Put(protected, "/sessions/{id}", sessionController.UpdateProposal,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Update a session proposal"),
		option.Summary("Update Session Proposal"),
		option.Tags("Session Proposal"),
	)

	fuego.Delete(protected, "/sessions/{id}", sessionController.DeleteProposal,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Delete a session proposal"),
		option.Summary("Delete Session Proposal"),
		option.Tags("Session Proposal"),
	)

	// Session routes
	fuego.Get(protected, "/sessions", sessionController.ListSessions,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("List all conference sessions"),
		option.Summary("List Sessions"),
		option.Tags("Sessions"),
	)

	fuego.Get(protected, "/sessions/{id}", sessionController.GetSession,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get session details by ID"),
		option.Summary("Get Session"),
		option.Tags("Sessions"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Post(protected, "/sessions/{id}/register", sessionController.RegisterForSession,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Register for a session"),
		option.Summary("Register for Session"),
		option.Tags("Sessions"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Post(protected, "/sessions/{id}/feedback", sessionController.CreateFeedback,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Leave feedback for a session"),
		option.Summary("Create Feedback"),
		option.Tags("Sessions"),
		option.Path("id", "Session ID", param.Required()),
	)

	// Coordinator routes
	coordinator := fuego.Group(v1, "/coordinator")
	fuego.Use(coordinator, middleware.AuthMiddleware(queries), middleware.RequireRole("event_coordinator"))

	fuego.Get(coordinator, "/proposals", coordinatorController.ListProposals,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("List all session proposals"),
		option.Summary("List Proposals"),
	)

	fuego.Put(coordinator, "/proposals/{id}/status", coordinatorController.UpdateProposalStatus,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Update proposal status (accept/reject)"),
		option.Summary("Update Proposal Status"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Delete(coordinator, "/sessions/{id}", coordinatorController.RemoveSession,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Remove a session"),
		option.Summary("Remove Session"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Delete(coordinator, "/feedback/{id}", coordinatorController.RemoveFeedback,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Remove inappropriate feedback"),
		option.Summary("Remove Feedback"),
		option.Path("id", "Feedback ID", param.Required()),
	)

	s.Run()
}
