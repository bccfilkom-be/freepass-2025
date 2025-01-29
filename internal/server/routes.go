package server

import (
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/server/middleware"
)

func (s *Server) registerRoutes() {
	// Root route
	fuego.Get(s.server, "/", func(c fuego.ContextNoBody) (fuego.HTML, error) {
		return "Conflux API is running! Check out <a href='http://localhost:8080/swagger/index.html'>the API Docs!</a>", nil
	})

	// API v1 routes
	v1 := fuego.Group(s.server, "/v1")
	s.registerAuthRoutes(v1)

	// Protected routes
	protected := fuego.Group(v1, "")
	fuego.Use(protected, middleware.AuthMiddleware(db.New(s.dbPool)))

	s.registerUserRoutes(protected)
	s.registerSessionRoutes(protected)
	s.registerCoordinatorRoutes(v1)
	s.registerAdminRoutes(v1)
}

func (s *Server) registerAuthRoutes(router *fuego.Server) {
	fuego.Post(router, "/users", s.controller.user.CreateUser, option.Tags("Auth"))
	fuego.Get(router, "/verify-email",
		s.controller.user.VerifyEmail,
		option.Query("token", "Token for email verification", param.Required()),
		option.Query("email", "Email for email verification", param.Required()),
		option.Tags("Auth"),
	)
	fuego.Post(router, "/login", s.controller.user.Login, option.Tags("Auth"))
}

func (s *Server) registerUserRoutes(router *fuego.Server) {
	fuego.Put(router, "/profile", s.controller.user.UpdateProfile,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Update user profile"),
		option.Summary("Update Profile"),
		option.Tags("User"),
	)

	fuego.Get(router, "/profile", s.controller.user.GetProfile,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get current user's profile"),
		option.Summary("Get Profile"),
		option.Tags("User"),
	)

	fuego.Get(router, "/users/{id}", s.controller.user.GetUserProfile,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get user profile by ID"),
		option.Summary("Get User Profile"),
		option.Tags("User"),
		option.Path("id", "User ID", param.Required()),
	)
}

func (s *Server) registerSessionRoutes(router *fuego.Server) {
	// Session proposal routes
	s.registerProposalRoutes(router)

	// Session management routes
	s.registerSessionManagementRoutes(router)
}

func (s *Server) registerProposalRoutes(router *fuego.Server) {
	fuego.Post(router, "/sessions", s.controller.session.CreateProposal,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Create a new session proposal"),
		option.Summary("Create Session Proposal"),
		option.Tags("Session Proposal"),
	)

	fuego.Put(router, "/sessions/{id}", s.controller.session.UpdateProposal,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Update a session proposal"),
		option.Summary("Update Session Proposal"),
		option.Tags("Session Proposal"),
	)

	fuego.Delete(router, "/sessions/{id}", s.controller.session.DeleteProposal,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Delete a session proposal"),
		option.Summary("Delete Session Proposal"),
		option.Tags("Session Proposal"),
	)

	fuego.Get(router, "/proposals", s.controller.session.GetUserProposals,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get user's submitted session proposals"),
		option.Summary("Get User Proposals"),
		option.Tags("Session Proposal"),
	)
}

func (s *Server) registerSessionManagementRoutes(router *fuego.Server) {
	fuego.Get(router, "/sessions", s.controller.session.ListSessions,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("List all conference sessions"),
		option.Summary("List Sessions"),
		option.Tags("Sessions"),
	)

	fuego.Get(router, "/sessions/{id}", s.controller.session.GetSession,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get session details by ID"),
		option.Summary("Get Session"),
		option.Tags("Sessions"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Post(router, "/sessions/{id}/register", s.controller.session.RegisterForSession,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Register for a session"),
		option.Summary("Register for Session"),
		option.Tags("Sessions"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Post(router, "/sessions/{id}/feedback", s.controller.session.CreateFeedback,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Leave feedback for a session"),
		option.Summary("Create Feedback"),
		option.Tags("Sessions"),
		option.Path("id", "Session ID", param.Required()),
	)
}

func (s *Server) registerCoordinatorRoutes(router *fuego.Server) {
	coordinator := fuego.Group(router, "/coordinator")
	fuego.Use(coordinator, middleware.AuthMiddleware(db.New(s.dbPool)), middleware.RequireRole("event_coordinator"))

	fuego.Get(coordinator, "/proposals", s.controller.coordinator.ListProposals,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("List all session proposals"),
		option.Summary("List Proposals"),
	)

	fuego.Put(coordinator, "/proposals/{id}/status", s.controller.coordinator.UpdateProposalStatus,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Update proposal status (accept/reject)"),
		option.Summary("Update Proposal Status"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Delete(coordinator, "/sessions/{id}", s.controller.coordinator.RemoveSession,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Remove a session"),
		option.Summary("Remove Session"),
		option.Path("id", "Session ID", param.Required()),
	)

	fuego.Delete(coordinator, "/feedback/{id}", s.controller.coordinator.RemoveFeedback,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Remove inappropriate feedback"),
		option.Summary("Remove Feedback"),
		option.Path("id", "Feedback ID", param.Required()),
	)
}

func (s *Server) registerAdminRoutes(router *fuego.Server) {
	admin := fuego.Group(router, "/admin")
	fuego.Use(admin, middleware.AuthMiddleware(db.New(s.dbPool)), middleware.RequireRole("admin"))

	fuego.Get(admin, "/users", s.controller.admin.GetAllUsers,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Get all users in the system"),
		option.Summary("Get All Users"),
		option.Tags("Admin"),
	)

	fuego.Put(admin, "/users/{id}/role", s.controller.admin.UpdateUserRole,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Update user role (promote/demote)"),
		option.Summary("Update User Role"),
		option.Path("id", "User ID", param.Required()),
	)

	fuego.Delete(admin, "/users/{id}", s.controller.admin.DeleteUser,
		option.Header("Authorization", "Bearer <token>", param.Required()),
		option.Description("Remove a user from the system"),
		option.Summary("Delete User"),
		option.Path("id", "User ID", param.Required()),
	)
}
