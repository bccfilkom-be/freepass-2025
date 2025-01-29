package model

// RoleUpdate represents the request body for updating a user's role
type RoleUpdate struct {
	Role string `json:"role" validate:"required,oneof=user event_coordinator"`
}
