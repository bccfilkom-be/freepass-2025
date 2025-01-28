package config

// Implementation of custom errors

var (
	// User errors
	ErrUserNotFound = "user not found"
	ErrInvalidEmailOrPassword = "invalid email or password"
	ErrEmailNotVerified = "email is not verified"

	// Session errors
	ErrPendingProposalExists    = "user already has a pending proposal"
	ErrProposalsNotOpen         = "session proposals are not currently open"
	ErrNoConferenceConfig       = "conference configuration not found - proposals are not open"
	ErrUnauthorizedSession      = "unauthorized to update this session"
	ErrOnlyPendingUpdate        = "can only update pending sessions"
	ErrUnauthorizedSessionDel   = "unauthorized to delete this session"
	ErrOnlyPendingDelete        = "can only delete pending sessions"
	ErrSessionNotFound          = "session not found"
)
