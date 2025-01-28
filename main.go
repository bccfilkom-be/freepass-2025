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

	// Initialize controllers
	userController := controller.NewUserController(userService)

	// Initialize router with custom error handler
	s := fuego.NewServer(
		fuego.WithAddr(os.Getenv("APP_URL")),
	)

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (fuego.HTML, error) {
		return "Conflux API is running! Check out <a href='http://localhost:8080/swagger/index.html'>the API Docs!</a>", nil
	})

	v1 := fuego.Group(s, "/v1")

	fuego.Post(v1, "/users", userController.CreateUser)
	fuego.Get(v1, "/verify-email",
		userController.VerifyEmail,
		option.Query("token", "Token for email verification", param.Required()),
		option.Query("email", "Email for email verification", param.Required()),
	)
	fuego.Post(v1, "/login", userController.Login)

	// Protected routes group
	protected := fuego.Group(v1, "")
	fuego.Use(protected, middleware.AuthMiddleware)

	fuego.Put(protected, "/profile", userController.UpdateProfile,
		option.Header("Authorization", "Bearer token", param.Required()),
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

	s.Run()
}
