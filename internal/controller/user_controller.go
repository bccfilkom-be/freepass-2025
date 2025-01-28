package controller

import (
	"errors"

	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/litegral/freepass-2025/internal/model"
	"github.com/litegral/freepass-2025/internal/service"
)

// UserController handles user-related operations
type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// CreateUser handles user creation request
func (c *UserController) CreateUser(ctx fuego.ContextWithBody[model.UserCreate]) (model.User, error) {
	body, err := ctx.Body()
	if err != nil {
		return model.User{}, err
	}

	user, err := c.userService.CreateUser(ctx.Context(), body)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return model.User{}, fuego.ConflictError{Title: "Email already exists"}
		}
		return model.User{}, err
	}

	return user, nil
}

// VerifyEmail handles email verification request
func (c *UserController) VerifyEmail(ctx fuego.ContextNoBody) (fuego.HTML, error) {
	token := ctx.QueryParam("token")
	email := ctx.QueryParam("email")

	if token == "" || email == "" {
		return "", fuego.BadRequestError{Title: "Token and email are required"}
	}

	err := c.userService.VerifyEmail(ctx.Context(), token, email)
	if err != nil {
		return "", fuego.BadRequestError{Title: "Invalid token or email"}
	}

	return fuego.HTML("Email verified successfully! You can now login to your account."), nil
}

// Login handles user login request
func (c *UserController) Login(ctx fuego.ContextWithBody[model.UserLogin]) (model.UserLoginResponse, error) {
	body, err := ctx.Body()
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	response, err := c.userService.Login(ctx.Context(), body)
	if err != nil {
		return model.UserLoginResponse{}, fuego.UnauthorizedError{Title: err.Error()}
	}

	return response, nil
}
