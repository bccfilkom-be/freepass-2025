package controller

import (
	"strconv"

	"github.com/go-fuego/fuego"
	"github.com/litegral/freepass-2025/internal/lib/config"
	"github.com/litegral/freepass-2025/internal/lib/jwt"
	"github.com/litegral/freepass-2025/internal/model"
	"github.com/litegral/freepass-2025/internal/service"
)

// AdminController handles administrator endpoints
type AdminController struct {
	adminService *service.AdminService
}

func NewAdminController(adminService *service.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

// UpdateUserRole handles updating a user's role
func (c *AdminController) UpdateUserRole(ctx fuego.ContextWithBody[model.RoleUpdate]) (model.User, error) {
	// Get user ID from path and convert to int
	userID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return model.User{}, fuego.BadRequestError{Title: "Invalid user ID"}
	}

	// Get request body
	update, err := ctx.Body()
	if err != nil {
		return model.User{}, err
	}

	// Update user role
	user, err := c.adminService.UpdateUserRole(ctx.Context(), int32(userID), update.Role)
	if err != nil {
		if err.Error() == config.ErrUserNotFound {
			return model.User{}, fuego.NotFoundError{Title: err.Error()}
		}
		return model.User{}, fuego.BadRequestError{Title: err.Error()}
	}

	return user, nil
}

// DeleteUser handles removing a user from the system
func (c *AdminController) DeleteUser(ctx fuego.ContextNoBody) (any, error) {
	// Get user ID from path and convert to int
	userID, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{Title: "Invalid user ID"}
	}

	// Delete user
	err = c.adminService.DeleteUser(ctx.Context(), int32(userID))
	if err != nil {
		if err.Error() == config.ErrUserNotFound {
			return nil, fuego.NotFoundError{Title: err.Error()}
		}
		return nil, fuego.InternalServerError{Title: err.Error()}
	}

	return map[string]string{"message": "User deleted successfully"}, nil
}

// GetAllUsers handles retrieving all users in the system
func (c *AdminController) GetAllUsers(ctx fuego.ContextNoBody) ([]model.User, error) {
	// Get current user ID
	claims, ok := ctx.Value(jwt.ClaimsContextKey).(*jwt.Claims)
	if !ok {
		return nil, fuego.UnauthorizedError{Title: "Invalid token claims"}
	}

	users, err := c.adminService.GetAllUsers(ctx.Context(), int32(claims.UserID))
	if err != nil {
		return nil, fuego.InternalServerError{Title: err.Error()}
	}

	return users, nil
}
