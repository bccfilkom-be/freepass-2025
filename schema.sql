CREATE TYPE user_role AS ENUM ('user', 'event_coordinator', 'admin');
CREATE TYPE session_status AS ENUM ('pending', 'accepted', 'rejected');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,  
    role user_role NOT NULL DEFAULT 'user',
    full_name TEXT,
    profile_pict_url TEXT,
    affiliation TEXT,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    verified_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE email_verification_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    token UUID NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL
);

-- Conference Sessions
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    room TEXT,  
    status session_status NOT NULL DEFAULT 'pending',
    seating_capacity INT NOT NULL,
    proposer_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CHECK (end_time > start_time)
);

CREATE TABLE session_registrations (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    session_id INT NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    registered_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (user_id, session_id)
);

-- Session Feedback
CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    session_id INT NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE conference_config (
    id SERIAL PRIMARY KEY,
    registration_start TIMESTAMPTZ NOT NULL,
    registration_end TIMESTAMPTZ NOT NULL,
    session_proposal_start TIMESTAMPTZ NOT NULL,
    session_proposal_end TIMESTAMPTZ NOT NULL,
    CHECK (registration_end > registration_start),
    CHECK (session_proposal_end > session_proposal_start)
);

-- Indexes
CREATE INDEX idx_sessions_time ON sessions (start_time, end_time);
CREATE INDEX idx_registrations_user ON session_registrations (user_id);
CREATE INDEX idx_feedback_session ON feedback (session_id);
CREATE INDEX idx_verification_token ON email_verification_tokens (token);
CREATE INDEX idx_verification_user ON email_verification_tokens (user_id);

-- Functions
-- Check if a user is already registered for an overlapping session
CREATE OR REPLACE FUNCTION check_session_overlap()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM session_registrations sr
        JOIN sessions s ON sr.session_id = s.id
        JOIN sessions new_s ON NEW.session_id = new_s.id
        WHERE sr.user_id = NEW.user_id
        AND s.id != new_s.id
        AND (new_s.start_time, new_s.end_time) OVERLAPS (s.start_time, s.end_time)
    ) THEN
        RAISE EXCEPTION 'Cannot register for overlapping sessions';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Check if a user has already submitted a proposal for the current conference period
CREATE OR REPLACE FUNCTION check_proposal_limit()
RETURNS TRIGGER AS $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM sessions
        WHERE proposer_id = NEW.proposer_id
        AND created_at >= (SELECT session_proposal_start FROM conference_config)
        AND created_at <= (SELECT session_proposal_end FROM conference_config)
    ) THEN
        RAISE EXCEPTION 'User can only submit one proposal per conference period';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Prevent creation of verified tokens for already verified users
CREATE OR REPLACE FUNCTION prevent_verified_tokens()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT is_verified FROM users WHERE id = NEW.user_id) THEN
        RAISE EXCEPTION 'User already verified';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Set the expiration time for a token to 24 hours after creation
CREATE OR REPLACE FUNCTION set_token_expiration()
RETURNS TRIGGER AS $$
BEGIN
    NEW.expires_at := NEW.created_at + INTERVAL '24 hours';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger Bindings
CREATE TRIGGER prevent_overlapping_registrations
BEFORE INSERT ON session_registrations
FOR EACH ROW EXECUTE FUNCTION check_session_overlap();

CREATE TRIGGER enforce_proposal_limit
BEFORE INSERT ON sessions
FOR EACH ROW EXECUTE FUNCTION check_proposal_limit();

CREATE TRIGGER verify_token_prevention
BEFORE INSERT ON email_verification_tokens
FOR EACH ROW EXECUTE FUNCTION prevent_verified_tokens();

CREATE TRIGGER set_token_expiration_trigger
BEFORE INSERT ON email_verification_tokens
FOR EACH ROW EXECUTE FUNCTION set_token_expiration();