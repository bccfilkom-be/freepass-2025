-- User Management
-- name: CreateUser :one
INSERT INTO users (email, password_hash, full_name, profile_pict_url, affiliation)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET full_name = $2, affiliation = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- Session Management
-- name: CreateSessionProposal :one
INSERT INTO sessions (
    title, description, start_time, end_time, 
    room, seating_capacity, proposer_id
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateSession :one
UPDATE sessions
SET title = $2, description = $3, start_time = $4,
    end_time = $5, room = $6, seating_capacity = $7,
    status = $8, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetSessionByID :one
SELECT s.*, u.full_name AS proposer_name, u.affiliation AS proposer_affiliation
FROM sessions s
JOIN users u ON s.proposer_id = u.id
WHERE s.id = $1;

-- name: ListSessions :many
SELECT * FROM sessions
ORDER BY start_time;

-- name: SoftDeleteSession :exec
UPDATE sessions
SET is_deleted = TRUE
WHERE id = $1;

-- Registration Management
-- name: RegisterForSession :one
INSERT INTO session_registrations (user_id, session_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetRegistration :one
SELECT * FROM session_registrations
WHERE user_id = $1 AND session_id = $2;

-- name: ListUserRegistrations :many
SELECT s.* FROM sessions s
JOIN session_registrations sr ON s.id = sr.session_id
WHERE sr.user_id = $1;

-- name: CountSessionRegistrations :one
SELECT COUNT(*) FROM session_registrations
WHERE session_id = $1;

-- Feedback Management
-- name: CreateFeedback :one
INSERT INTO feedback (user_id, session_id, comment)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListSessionFeedback :many
SELECT f.*, u.full_name, u.affiliation, u.profile_pict_url
FROM feedback f
JOIN users u ON f.user_id = u.id
WHERE f.session_id = $1 AND f.is_deleted = FALSE;

-- name: SoftDeleteFeedback :exec
UPDATE feedback
SET is_deleted = TRUE
WHERE id = $1;

-- Conference Configuration
-- name: GetConferenceConfig :one
SELECT * FROM conference_config
ORDER BY id DESC
LIMIT 1;

-- name: UpdateConferenceConfig :one
INSERT INTO conference_config (
    registration_start, registration_end,
    session_proposal_start, session_proposal_end
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- Email Verification
-- name: CreateVerificationToken :one
INSERT INTO email_verification_tokens (user_id, token)
VALUES ($1, $2)
RETURNING *;

-- name: VerifyEmail :one
WITH deleted_token AS (
    DELETE FROM email_verification_tokens
    WHERE token = $1 AND expires_at > NOW()
    RETURNING user_id
)
UPDATE users
SET is_verified = TRUE, verified_at = NOW()
FROM deleted_token
WHERE users.id = deleted_token.user_id
RETURNING users.*;

-- name: GetVerificationToken :one
SELECT * FROM email_verification_tokens
WHERE token = $1;

-- Admin Operations
-- name: UpdateUserRole :one
UPDATE users
SET role = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- Coordinator Operations
-- name: ListPendingProposals :many
SELECT * FROM sessions
WHERE status = 'pending'
ORDER BY created_at;

-- name: UpdateSessionStatus :one
UPDATE sessions
SET status = $2
WHERE id = $1
RETURNING *;

-- name: ListSessionProposals :many
SELECT s.*, u.full_name AS proposer_name, u.affiliation AS proposer_affiliation
FROM sessions s
JOIN users u ON s.proposer_id = u.id
WHERE s.status = 'pending' AND s.is_deleted = FALSE
ORDER BY s.created_at DESC;

-- name: UpdateSessionStatusByCoordinator :one
UPDATE sessions
SET status = $2, updated_at = NOW()
WHERE id = $1 AND status = 'pending' AND is_deleted = FALSE
RETURNING *;

-- name: SoftDeleteSessionByCoordinator :exec
UPDATE sessions
SET is_deleted = TRUE, updated_at = NOW()
WHERE id = $1 AND is_deleted = FALSE;

-- name: SoftDeleteFeedbackByCoordinator :exec
UPDATE feedback
SET is_deleted = TRUE
WHERE id = $1 AND is_deleted = FALSE;