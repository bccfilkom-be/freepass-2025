package dto

import "time"

type GetUserDetailResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserProfileResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserProfileRequest struct {
	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio"`
}

type UpdateUserRoleRequest struct {
	Role string `json:"role" validate:"required,oneof=ADMIN COORDINATOR USER"`
}
