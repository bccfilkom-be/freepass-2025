// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countSessionRegistrations = `-- name: CountSessionRegistrations :one
SELECT COUNT(*) FROM session_registrations
WHERE session_id = $1
`

func (q *Queries) CountSessionRegistrations(ctx context.Context, sessionID int32) (int64, error) {
	row := q.db.QueryRow(ctx, countSessionRegistrations, sessionID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createFeedback = `-- name: CreateFeedback :one
INSERT INTO feedback (user_id, session_id, comment)
VALUES ($1, $2, $3)
RETURNING id, user_id, session_id, comment, created_at, is_deleted
`

type CreateFeedbackParams struct {
	UserID    int32
	SessionID int32
	Comment   string
}

// Feedback Management
func (q *Queries) CreateFeedback(ctx context.Context, arg CreateFeedbackParams) (Feedback, error) {
	row := q.db.QueryRow(ctx, createFeedback, arg.UserID, arg.SessionID, arg.Comment)
	var i Feedback
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SessionID,
		&i.Comment,
		&i.CreatedAt,
		&i.IsDeleted,
	)
	return i, err
}

const createSessionProposal = `-- name: CreateSessionProposal :one
INSERT INTO sessions (
    title, description, start_time, end_time, 
    room, seating_capacity, proposer_id
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, title, description, start_time, end_time, room, status, seating_capacity, proposer_id, created_at, updated_at, is_deleted
`

type CreateSessionProposalParams struct {
	Title           string
	Description     pgtype.Text
	StartTime       pgtype.Timestamptz
	EndTime         pgtype.Timestamptz
	Room            pgtype.Text
	SeatingCapacity int32
	ProposerID      pgtype.Int4
}

// Session Management
func (q *Queries) CreateSessionProposal(ctx context.Context, arg CreateSessionProposalParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSessionProposal,
		arg.Title,
		arg.Description,
		arg.StartTime,
		arg.EndTime,
		arg.Room,
		arg.SeatingCapacity,
		arg.ProposerID,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.StartTime,
		&i.EndTime,
		&i.Room,
		&i.Status,
		&i.SeatingCapacity,
		&i.ProposerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password_hash, full_name, profile_pict_url, affiliation)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, email, password_hash, role, full_name, profile_pict_url, affiliation, is_verified, verified_at, created_at, updated_at
`

type CreateUserParams struct {
	Email          string
	PasswordHash   string
	FullName       pgtype.Text
	ProfilePictUrl pgtype.Text
	Affiliation    pgtype.Text
}

// User Management
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Email,
		arg.PasswordHash,
		arg.FullName,
		arg.ProfilePictUrl,
		arg.Affiliation,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.FullName,
		&i.ProfilePictUrl,
		&i.Affiliation,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createVerificationToken = `-- name: CreateVerificationToken :one
INSERT INTO email_verification_tokens (user_id, token)
VALUES ($1, $2)
RETURNING id, user_id, token, created_at, expires_at
`

type CreateVerificationTokenParams struct {
	UserID int32
	Token  pgtype.UUID
}

// Email Verification
func (q *Queries) CreateVerificationToken(ctx context.Context, arg CreateVerificationTokenParams) (EmailVerificationToken, error) {
	row := q.db.QueryRow(ctx, createVerificationToken, arg.UserID, arg.Token)
	var i EmailVerificationToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, email, full_name, profile_pict_url, affiliation, role, is_verified, verified_at, created_at, updated_at FROM users
`

type GetAllUsersRow struct {
	ID             int32
	Email          string
	FullName       pgtype.Text
	ProfilePictUrl pgtype.Text
	Affiliation    pgtype.Text
	Role           UserRole
	IsVerified     bool
	VerifiedAt     pgtype.Timestamptz
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
}

// Admin Operations
func (q *Queries) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersRow
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FullName,
			&i.ProfilePictUrl,
			&i.Affiliation,
			&i.Role,
			&i.IsVerified,
			&i.VerifiedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getConferenceConfig = `-- name: GetConferenceConfig :one
SELECT id, registration_start, registration_end, session_proposal_start, session_proposal_end FROM conference_config
ORDER BY id DESC
LIMIT 1
`

// Conference Configuration
func (q *Queries) GetConferenceConfig(ctx context.Context) (ConferenceConfig, error) {
	row := q.db.QueryRow(ctx, getConferenceConfig)
	var i ConferenceConfig
	err := row.Scan(
		&i.ID,
		&i.RegistrationStart,
		&i.RegistrationEnd,
		&i.SessionProposalStart,
		&i.SessionProposalEnd,
	)
	return i, err
}

const getRegistration = `-- name: GetRegistration :one
SELECT id, user_id, session_id, registered_at FROM session_registrations
WHERE user_id = $1 AND session_id = $2
`

type GetRegistrationParams struct {
	UserID    int32
	SessionID int32
}

func (q *Queries) GetRegistration(ctx context.Context, arg GetRegistrationParams) (SessionRegistration, error) {
	row := q.db.QueryRow(ctx, getRegistration, arg.UserID, arg.SessionID)
	var i SessionRegistration
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SessionID,
		&i.RegisteredAt,
	)
	return i, err
}

const getSessionByID = `-- name: GetSessionByID :one
SELECT s.id, s.title, s.description, s.start_time, s.end_time, s.room, s.status, s.seating_capacity, s.proposer_id, s.created_at, s.updated_at, s.is_deleted, u.full_name AS proposer_name, u.affiliation AS proposer_affiliation
FROM sessions s
JOIN users u ON s.proposer_id = u.id
WHERE s.id = $1 AND s.is_deleted = FALSE
`

type GetSessionByIDRow struct {
	ID                  int32
	Title               string
	Description         pgtype.Text
	StartTime           pgtype.Timestamptz
	EndTime             pgtype.Timestamptz
	Room                pgtype.Text
	Status              SessionStatus
	SeatingCapacity     int32
	ProposerID          pgtype.Int4
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
	IsDeleted           pgtype.Bool
	ProposerName        pgtype.Text
	ProposerAffiliation pgtype.Text
}

func (q *Queries) GetSessionByID(ctx context.Context, id int32) (GetSessionByIDRow, error) {
	row := q.db.QueryRow(ctx, getSessionByID, id)
	var i GetSessionByIDRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.StartTime,
		&i.EndTime,
		&i.Room,
		&i.Status,
		&i.SeatingCapacity,
		&i.ProposerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
		&i.ProposerName,
		&i.ProposerAffiliation,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password_hash, role, full_name, profile_pict_url, affiliation, is_verified, verified_at, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.FullName,
		&i.ProfilePictUrl,
		&i.Affiliation,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, password_hash, role, full_name, profile_pict_url, affiliation, is_verified, verified_at, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.FullName,
		&i.ProfilePictUrl,
		&i.Affiliation,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getVerificationToken = `-- name: GetVerificationToken :one
SELECT id, user_id, token, created_at, expires_at FROM email_verification_tokens
WHERE token = $1
`

func (q *Queries) GetVerificationToken(ctx context.Context, token pgtype.UUID) (EmailVerificationToken, error) {
	row := q.db.QueryRow(ctx, getVerificationToken, token)
	var i EmailVerificationToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Token,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const listPendingProposals = `-- name: ListPendingProposals :many
SELECT id, title, description, start_time, end_time, room, status, seating_capacity, proposer_id, created_at, updated_at, is_deleted FROM sessions
WHERE status = 'pending'
ORDER BY created_at
`

// Coordinator Operations
func (q *Queries) ListPendingProposals(ctx context.Context) ([]Session, error) {
	rows, err := q.db.Query(ctx, listPendingProposals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.StartTime,
			&i.EndTime,
			&i.Room,
			&i.Status,
			&i.SeatingCapacity,
			&i.ProposerID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessionFeedback = `-- name: ListSessionFeedback :many
SELECT f.id, f.user_id, f.session_id, f.comment, f.created_at, f.is_deleted, u.full_name, u.affiliation, u.profile_pict_url
FROM feedback f
JOIN users u ON f.user_id = u.id
WHERE f.session_id = $1 AND f.is_deleted = FALSE
`

type ListSessionFeedbackRow struct {
	ID             int32
	UserID         int32
	SessionID      int32
	Comment        string
	CreatedAt      pgtype.Timestamptz
	IsDeleted      pgtype.Bool
	FullName       pgtype.Text
	Affiliation    pgtype.Text
	ProfilePictUrl pgtype.Text
}

func (q *Queries) ListSessionFeedback(ctx context.Context, sessionID int32) ([]ListSessionFeedbackRow, error) {
	rows, err := q.db.Query(ctx, listSessionFeedback, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSessionFeedbackRow
	for rows.Next() {
		var i ListSessionFeedbackRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.SessionID,
			&i.Comment,
			&i.CreatedAt,
			&i.IsDeleted,
			&i.FullName,
			&i.Affiliation,
			&i.ProfilePictUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessionProposals = `-- name: ListSessionProposals :many
SELECT s.id, s.title, s.description, s.start_time, s.end_time, s.room, s.status, s.seating_capacity, s.proposer_id, s.created_at, s.updated_at, s.is_deleted, u.full_name AS proposer_name, u.affiliation AS proposer_affiliation
FROM sessions s
JOIN users u ON s.proposer_id = u.id
WHERE s.status = 'pending' AND s.is_deleted = FALSE
ORDER BY s.created_at DESC
`

type ListSessionProposalsRow struct {
	ID                  int32
	Title               string
	Description         pgtype.Text
	StartTime           pgtype.Timestamptz
	EndTime             pgtype.Timestamptz
	Room                pgtype.Text
	Status              SessionStatus
	SeatingCapacity     int32
	ProposerID          pgtype.Int4
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
	IsDeleted           pgtype.Bool
	ProposerName        pgtype.Text
	ProposerAffiliation pgtype.Text
}

func (q *Queries) ListSessionProposals(ctx context.Context) ([]ListSessionProposalsRow, error) {
	rows, err := q.db.Query(ctx, listSessionProposals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSessionProposalsRow
	for rows.Next() {
		var i ListSessionProposalsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.StartTime,
			&i.EndTime,
			&i.Room,
			&i.Status,
			&i.SeatingCapacity,
			&i.ProposerID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
			&i.ProposerName,
			&i.ProposerAffiliation,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessions = `-- name: ListSessions :many
SELECT id, title, description, start_time, end_time, room, status, seating_capacity, proposer_id, created_at, updated_at, is_deleted FROM sessions
WHERE is_deleted = FALSE
ORDER BY start_time
`

func (q *Queries) ListSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.db.Query(ctx, listSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.StartTime,
			&i.EndTime,
			&i.Room,
			&i.Status,
			&i.SeatingCapacity,
			&i.ProposerID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserRegistrations = `-- name: ListUserRegistrations :many
SELECT s.id, s.title, s.description, s.start_time, s.end_time, s.room, s.status, s.seating_capacity, s.proposer_id, s.created_at, s.updated_at, s.is_deleted FROM sessions s
JOIN session_registrations sr ON s.id = sr.session_id
WHERE sr.user_id = $1
`

func (q *Queries) ListUserRegistrations(ctx context.Context, userID int32) ([]Session, error) {
	rows, err := q.db.Query(ctx, listUserRegistrations, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.StartTime,
			&i.EndTime,
			&i.Room,
			&i.Status,
			&i.SeatingCapacity,
			&i.ProposerID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const registerForSession = `-- name: RegisterForSession :one
INSERT INTO session_registrations (user_id, session_id)
VALUES ($1, $2)
RETURNING id, user_id, session_id, registered_at
`

type RegisterForSessionParams struct {
	UserID    int32
	SessionID int32
}

// Registration Management
func (q *Queries) RegisterForSession(ctx context.Context, arg RegisterForSessionParams) (SessionRegistration, error) {
	row := q.db.QueryRow(ctx, registerForSession, arg.UserID, arg.SessionID)
	var i SessionRegistration
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.SessionID,
		&i.RegisteredAt,
	)
	return i, err
}

const softDeleteFeedback = `-- name: SoftDeleteFeedback :exec
UPDATE feedback
SET is_deleted = TRUE
WHERE id = $1
`

func (q *Queries) SoftDeleteFeedback(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, softDeleteFeedback, id)
	return err
}

const softDeleteFeedbackByCoordinator = `-- name: SoftDeleteFeedbackByCoordinator :exec
UPDATE feedback
SET is_deleted = TRUE
WHERE id = $1 AND is_deleted = FALSE
`

func (q *Queries) SoftDeleteFeedbackByCoordinator(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, softDeleteFeedbackByCoordinator, id)
	return err
}

const softDeleteSession = `-- name: SoftDeleteSession :exec
UPDATE sessions
SET is_deleted = TRUE
WHERE id = $1
`

func (q *Queries) SoftDeleteSession(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, softDeleteSession, id)
	return err
}

const softDeleteSessionByCoordinator = `-- name: SoftDeleteSessionByCoordinator :exec
UPDATE sessions
SET is_deleted = TRUE, updated_at = NOW()
WHERE id = $1 AND is_deleted = FALSE
`

func (q *Queries) SoftDeleteSessionByCoordinator(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, softDeleteSessionByCoordinator, id)
	return err
}

const updateConferenceConfig = `-- name: UpdateConferenceConfig :one
INSERT INTO conference_config (
    registration_start, registration_end,
    session_proposal_start, session_proposal_end
) VALUES ($1, $2, $3, $4)
RETURNING id, registration_start, registration_end, session_proposal_start, session_proposal_end
`

type UpdateConferenceConfigParams struct {
	RegistrationStart    pgtype.Timestamptz
	RegistrationEnd      pgtype.Timestamptz
	SessionProposalStart pgtype.Timestamptz
	SessionProposalEnd   pgtype.Timestamptz
}

func (q *Queries) UpdateConferenceConfig(ctx context.Context, arg UpdateConferenceConfigParams) (ConferenceConfig, error) {
	row := q.db.QueryRow(ctx, updateConferenceConfig,
		arg.RegistrationStart,
		arg.RegistrationEnd,
		arg.SessionProposalStart,
		arg.SessionProposalEnd,
	)
	var i ConferenceConfig
	err := row.Scan(
		&i.ID,
		&i.RegistrationStart,
		&i.RegistrationEnd,
		&i.SessionProposalStart,
		&i.SessionProposalEnd,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :one
UPDATE sessions
SET title = $2, description = $3, start_time = $4,
    end_time = $5, room = $6, seating_capacity = $7,
    status = $8, updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, start_time, end_time, room, status, seating_capacity, proposer_id, created_at, updated_at, is_deleted
`

type UpdateSessionParams struct {
	ID              int32
	Title           string
	Description     pgtype.Text
	StartTime       pgtype.Timestamptz
	EndTime         pgtype.Timestamptz
	Room            pgtype.Text
	SeatingCapacity int32
	Status          SessionStatus
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, updateSession,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.StartTime,
		arg.EndTime,
		arg.Room,
		arg.SeatingCapacity,
		arg.Status,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.StartTime,
		&i.EndTime,
		&i.Room,
		&i.Status,
		&i.SeatingCapacity,
		&i.ProposerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
	)
	return i, err
}

const updateSessionStatus = `-- name: UpdateSessionStatus :one
UPDATE sessions
SET status = $2
WHERE id = $1
RETURNING id, title, description, start_time, end_time, room, status, seating_capacity, proposer_id, created_at, updated_at, is_deleted
`

type UpdateSessionStatusParams struct {
	ID     int32
	Status SessionStatus
}

func (q *Queries) UpdateSessionStatus(ctx context.Context, arg UpdateSessionStatusParams) (Session, error) {
	row := q.db.QueryRow(ctx, updateSessionStatus, arg.ID, arg.Status)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.StartTime,
		&i.EndTime,
		&i.Room,
		&i.Status,
		&i.SeatingCapacity,
		&i.ProposerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
	)
	return i, err
}

const updateSessionStatusByCoordinator = `-- name: UpdateSessionStatusByCoordinator :one
UPDATE sessions
SET status = $2, updated_at = NOW()
WHERE id = $1 AND status = 'pending' AND is_deleted = FALSE
RETURNING id, title, description, start_time, end_time, room, status, seating_capacity, proposer_id, created_at, updated_at, is_deleted
`

type UpdateSessionStatusByCoordinatorParams struct {
	ID     int32
	Status SessionStatus
}

func (q *Queries) UpdateSessionStatusByCoordinator(ctx context.Context, arg UpdateSessionStatusByCoordinatorParams) (Session, error) {
	row := q.db.QueryRow(ctx, updateSessionStatusByCoordinator, arg.ID, arg.Status)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.StartTime,
		&i.EndTime,
		&i.Room,
		&i.Status,
		&i.SeatingCapacity,
		&i.ProposerID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.IsDeleted,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET full_name = $2, affiliation = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, email, password_hash, role, full_name, profile_pict_url, affiliation, is_verified, verified_at, created_at, updated_at
`

type UpdateUserParams struct {
	ID          int32
	FullName    pgtype.Text
	Affiliation pgtype.Text
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.ID, arg.FullName, arg.Affiliation)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.FullName,
		&i.ProfilePictUrl,
		&i.Affiliation,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE users
SET role = $2
WHERE id = $1
RETURNING id, email, password_hash, role, full_name, profile_pict_url, affiliation, is_verified, verified_at, created_at, updated_at
`

type UpdateUserRoleParams struct {
	ID   int32
	Role UserRole
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUserRole, arg.ID, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.FullName,
		&i.ProfilePictUrl,
		&i.Affiliation,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifyEmail = `-- name: VerifyEmail :one
WITH deleted_token AS (
    DELETE FROM email_verification_tokens
    WHERE token = $1 AND expires_at > NOW()
    RETURNING user_id
)
UPDATE users
SET is_verified = TRUE, verified_at = NOW()
FROM deleted_token
WHERE users.id = deleted_token.user_id
RETURNING users.id, users.email, users.password_hash, users.role, users.full_name, users.profile_pict_url, users.affiliation, users.is_verified, users.verified_at, users.created_at, users.updated_at
`

func (q *Queries) VerifyEmail(ctx context.Context, token pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, verifyEmail, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.Role,
		&i.FullName,
		&i.ProfilePictUrl,
		&i.Affiliation,
		&i.IsVerified,
		&i.VerifiedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
