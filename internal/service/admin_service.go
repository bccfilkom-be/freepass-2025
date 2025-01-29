package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/litegral/freepass-2025/internal/lib/config"
	"github.com/litegral/freepass-2025/internal/lib/db"
	"github.com/litegral/freepass-2025/internal/model"
)

// AdminService handles administrator operations
type AdminService struct {
	queries *db.Queries
}

func NewAdminService(queries *db.Queries) *AdminService {
	return &AdminService{
		queries: queries,
	}
}

// GetAllUsers returns all users in the system except the current user
func (s *AdminService) GetAllUsers(ctx context.Context, currentUserID int32) ([]model.User, error) {
	dbUsers, err := s.queries.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	modelUsers := make([]model.User, 0, len(dbUsers))
	for _, u := range dbUsers {
		if int32(u.ID) == currentUserID {
			continue
		}
		modelUsers = append(modelUsers, model.User{
			ID:          int(u.ID),
			Email:       u.Email,
			Role:        string(u.Role),
			FullName:    u.FullName.String,
			Affiliation: u.Affiliation.String,
			IsVerified:  u.IsVerified,
			VerifiedAt:  u.VerifiedAt.Time,
			CreatedAt:   u.CreatedAt.Time,
			UpdatedAt:   u.UpdatedAt.Time,
		})
	}
	return modelUsers, nil
}

// UpdateUserRole updates a user's role
func (s *AdminService) UpdateUserRole(ctx context.Context, userID int32, role string) (model.User, error) {
	// First check if user exists
	_, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, errors.New(config.ErrUserNotFound)
		}
		return model.User{}, err
	}

	// Update user role
	updatedUser, err := s.queries.UpdateUserRole(ctx, db.UpdateUserRoleParams{
		ID:   userID,
		Role: db.UserRole(role),
	})
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:          int(updatedUser.ID),
		Email:       updatedUser.Email,
		Role:        string(updatedUser.Role),
		FullName:    updatedUser.FullName.String,
		Affiliation: updatedUser.Affiliation.String,
		IsVerified:  updatedUser.IsVerified,
		VerifiedAt:  updatedUser.VerifiedAt.Time,
		CreatedAt:   updatedUser.CreatedAt.Time,
		UpdatedAt:   updatedUser.UpdatedAt.Time,
	}, nil
}

// DeleteUser permanently deletes a user from the system
func (s *AdminService) DeleteUser(ctx context.Context, userID int32) error {
	// First check if user exists
	_, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New(config.ErrUserNotFound)
		}
		return err
	}

	// Delete user
	return s.queries.DeleteUser(ctx, userID)
}
