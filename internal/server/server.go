package server

import (
	"context"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/litegral/freepass-2025/internal/controller"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/service"
)

type Server struct {
	server     *fuego.Server
	dbPool     *pgxpool.Pool
	controller *Controllers
}

type Controllers struct {
	user        *controller.UserController
	session     *controller.SessionController
	coordinator *controller.CoordinatorController
	admin       *controller.AdminController
}

// NewServer initializes a new server instance with all dependencies
func NewServer(dbURL string, serverAddr string) (*Server, error) {
	// Initialize database connection pool
	dbPool, err := initDB(dbURL)
	if err != nil {
		return nil, err
	}

	// Initialize services and controllers
	controllers, err := initControllers(dbPool)
	if err != nil {
		return nil, err
	}

	s := fuego.NewServer(
		fuego.WithAddr(serverAddr),
	)

	return &Server{
		server:     s,
		dbPool:     dbPool,
		controller: controllers,
	}, nil
}

func initDB(dbURL string) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}

	// Set pool configuration
	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute

	return pgxpool.NewWithConfig(context.Background(), poolConfig)
}

func initControllers(dbPool *pgxpool.Pool) (*Controllers, error) {
	// Initialize SQLC queries
	queries := db.New(dbPool)

	// Initialize services
	emailService, err := service.NewEmailService()
	if err != nil {
		return nil, err
	}

	userService := service.NewUserService(queries, emailService)
	sessionService := service.NewSessionService(queries)
	coordinatorService := service.NewCoordinatorService(queries)
	adminService := service.NewAdminService(queries)

	// Initialize controllers
	return &Controllers{
		user:        controller.NewUserController(userService),
		session:     controller.NewSessionController(sessionService),
		coordinator: controller.NewCoordinatorController(coordinatorService),
		admin:       controller.NewAdminController(adminService),
	}, nil
}

// Start starts the server
func (s *Server) Start() error {
	defer s.dbPool.Close()

	// Register routes
	s.registerRoutes()

	return s.server.Run()
}
