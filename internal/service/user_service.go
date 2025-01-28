package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/litegral/freepass-2025/internal/lib"
	"github.com/litegral/freepass-2025/internal/lib/db"
	c "github.com/litegral/freepass-2025/internal/lib/config"
	"github.com/litegral/freepass-2025/internal/lib/jwt"
	"github.com/litegral/freepass-2025/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// UserService is a service for user operations
type UserService struct {
	queries      *db.Queries
	emailService *EmailService
}

func NewUserService(queries *db.Queries, emailService *EmailService) *UserService {
	return &UserService{
		queries:      queries,
		emailService: emailService,
	}
}

// CreateUser handles user creation
func (s *UserService) CreateUser(ctx context.Context, params model.UserCreate) (model.User, error) {
	// Hash the password
	hashedPassword, err := lib.HashPassword(params.Password)
	if err != nil {
		return model.User{}, err
	}
	params.Password = hashedPassword

	// Dicebear default profile picture URL
	profilePictUrl := "https://api.dicebear.com/9.x/bottts-neutral/svg?seed=" + params.Email

	// Create user
	user, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Email:          params.Email,
		PasswordHash:   params.Password,
		FullName:       pgtype.Text{String: params.FullName, Valid: true},
		Affiliation:    pgtype.Text{String: params.Affiliation, Valid: true},
		ProfilePictUrl: pgtype.Text{String: profilePictUrl, Valid: true},
	})
	if err != nil {
		return model.User{}, err
	}

	// Generate and store verification token
	token, err := lib.GenerateVerificationToken()
	if err != nil {
		return model.User{}, err
	}

	// Store the token in the database
	_, err = s.queries.CreateVerificationToken(ctx, db.CreateVerificationTokenParams{
		UserID: user.ID,
		Token:  pgtype.UUID{Bytes: token, Valid: true},
	})
	if err != nil {
		return model.User{}, err
	}

	// Send verification email
	if err := s.emailService.SendEmailVerification(ctx, user.Email, token.String()); err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:          int(user.ID),
		Email:       user.Email,
		Role:        string(user.Role),
		FullName:    user.FullName.String,
		Affiliation: user.Affiliation.String,
		IsVerified:  user.IsVerified,
		VerifiedAt:  user.VerifiedAt.Time,
		CreatedAt:   user.CreatedAt.Time,
		UpdatedAt:   user.UpdatedAt.Time,
	}, nil
}

// VerifyEmail handles email verification
func (s *UserService) VerifyEmail(ctx context.Context, token string, email string) error {
	// Parse the token
	tokenUUID, err := lib.ParseUUID(token)
	if err != nil {
		return err
	}

	// Verify the email using the token
	user, err := s.queries.VerifyEmail(ctx, pgtype.UUID{
		Bytes: tokenUUID,
		Valid: true,
	})
	if err != nil {
		return err
	}

	// Check if the verified user matches the email
	if user.Email != email {
		return err
	}

	return nil
}

// Login authenticates a user and returns a JWT token
func (s *UserService) Login(ctx context.Context, params model.UserLogin) (model.UserLoginResponse, error) {
	// Get user by email
	user, err := s.queries.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return model.UserLoginResponse{}, errors.New(c.ErrInvalidEmailOrPassword)
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if err != nil {
		return model.UserLoginResponse{}, errors.New(c.ErrInvalidEmailOrPassword)
	}

	// Check if email is verified
	if !user.IsVerified {
		return model.UserLoginResponse{}, errors.New(c.ErrEmailNotVerified)
	}

	// Convert DB user to model user
	modelUser := model.User{
		ID:          int(user.ID),
		Email:       user.Email,
		Role:        string(user.Role),
		FullName:    user.FullName.String,
		Affiliation: user.Affiliation.String,
		IsVerified:  user.IsVerified,
		VerifiedAt:  user.VerifiedAt.Time,
		CreatedAt:   user.CreatedAt.Time,
		UpdatedAt:   user.UpdatedAt.Time,
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(modelUser)
	if err != nil {
		return model.UserLoginResponse{}, err
	}

	return model.UserLoginResponse{
		User:  modelUser,
		Token: token,
	}, nil
}

// UpdateProfile updates a user's profile
func (s *UserService) UpdateProfile(ctx context.Context, userID int32, params model.UserProfileUpdate) (model.User, error) {
	// Update user profile
	user, err := s.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:          userID,
		FullName:    pgtype.Text{String: params.FullName, Valid: true},
		Affiliation: pgtype.Text{String: params.Affiliation, Valid: true},
	})
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:          int(user.ID),
		Email:       user.Email,
		Role:        string(user.Role),
		FullName:    user.FullName.String,
		Affiliation: user.Affiliation.String,
		IsVerified:  user.IsVerified,
		VerifiedAt:  user.VerifiedAt.Time,
		CreatedAt:   user.CreatedAt.Time,
		UpdatedAt:   user.UpdatedAt.Time,
	}, nil
}

// GetProfile retrieves a user's profile by ID
func (s *UserService) GetProfile(ctx context.Context, userID int32) (model.User, error) {
	user, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New(c.ErrUserNotFound)
		}
		return model.User{}, err 
	}

	return model.User{
		ID:          int(user.ID),
		Email:       user.Email,
		Role:        string(user.Role),
		FullName:    user.FullName.String,
		Affiliation: user.Affiliation.String,
		IsVerified:  user.IsVerified,
		VerifiedAt:  user.VerifiedAt.Time,
		CreatedAt:   user.CreatedAt.Time,
		UpdatedAt:   user.UpdatedAt.Time,
	}, nil
}
