package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/nathakusuma/astungkara/domain/contract"
	"github.com/nathakusuma/astungkara/domain/dto"
	"github.com/nathakusuma/astungkara/domain/entity"
	"github.com/nathakusuma/astungkara/domain/enum"
	"github.com/nathakusuma/astungkara/domain/errorpkg"
	"github.com/nathakusuma/astungkara/internal/infra/env"
	"github.com/nathakusuma/astungkara/pkg/bcrypt"
	"github.com/nathakusuma/astungkara/pkg/jwt"
	"github.com/nathakusuma/astungkara/pkg/log"
	"github.com/nathakusuma/astungkara/pkg/mail"
	"github.com/nathakusuma/astungkara/pkg/randgen"
	"github.com/nathakusuma/astungkara/pkg/uuidpkg"
	"github.com/redis/go-redis/v9"
	"net/url"
	"strconv"
	"time"
)

type authService struct {
	repo    contract.IAuthRepository
	userSvc contract.IUserService
	bcrypt  bcrypt.IBcrypt
	jwt     jwt.IJwt
	mailer  mail.IMailer
	uuid    uuidpkg.IUUID
}

func NewAuthService(
	authRepo contract.IAuthRepository,
	userSvc contract.IUserService,
	bcrypt bcrypt.IBcrypt,
	jwt jwt.IJwt,
	mailer mail.IMailer,
	uuid uuidpkg.IUUID,
) contract.IAuthService {
	return &authService{
		repo:    authRepo,
		userSvc: userSvc,
		bcrypt:  bcrypt,
		jwt:     jwt,
		mailer:  mailer,
		uuid:    uuid,
	}
}

func (s *authService) RequestOTPRegisterUser(ctx context.Context, email string) error {
	// check if email is already registered
	_, err := s.userSvc.GetUserByEmail(ctx, email)
	if err == nil {
		return errorpkg.ErrEmailAlreadyRegistered
	}

	if !errors.Is(err, errorpkg.ErrNotFound) {
		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
			"email": email,
		}, "[AuthService][RequestOTPRegisterUser] failed to get user by email")

		return errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	// generate otp
	otp := strconv.Itoa(randgen.RandomNumber(6))

	// save otp
	err = s.repo.SetUserRegisterOTP(ctx, email, otp)
	if err != nil {
		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
			"email": email,
		}, "[AuthService][RequestOTPRegisterUser] failed to save otp")

		return errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	// send otp to email
	go func() {
		err = s.mailer.Send(
			email,
			"[Astungkara] Verify Your Account",
			"otp_register_user.html",
			map[string]interface{}{
				"otp":  otp,
				"href": env.GetEnv().FrontendURL + "/verify-email?email=" + url.QueryEscape(email) + "&otp=" + otp,
			})

		if err != nil {
			log.Error(map[string]interface{}{
				"error": err.Error(),
			}, "[AuthService][RequestOTPRegisterUser] failed to send email")
		}
	}()

	return nil
}

func (s *authService) CheckOTPRegisterUser(ctx context.Context, email, otp string) error {
	savedOtp, err := s.repo.GetUserRegisterOTP(ctx, email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return errorpkg.ErrInvalidOTP
		}

		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
			"email": email,
		}, "[AuthService][CheckOTPRegisterUser] failed to get otp")

		return errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	if savedOtp != otp {
		return errorpkg.ErrInvalidOTP
	}

	return nil
}

func (s *authService) RegisterUser(ctx context.Context,
	req dto.RegisterUserRequest) (dto.LoginResponse, error) {

	var resp dto.LoginResponse

	// req without Password and OTP
	loggableReq := req
	loggableReq.Password = ""
	loggableReq.OTP = ""

	// get otp
	savedOtp, err := s.repo.GetUserRegisterOTP(ctx, req.Email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return resp, errorpkg.ErrInvalidOTP
		}

		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
			"req":   loggableReq,
		}, "[AuthService][RegisterUser] failed to get otp")

		return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	if savedOtp != req.OTP {
		return resp, errorpkg.ErrInvalidOTP
	}

	// delete otp
	err = s.repo.DeleteUserRegisterOTP(ctx, req.Email)
	if err != nil {
		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
			"req":   loggableReq,
		}, "[AuthService][RegisterUser] failed to delete otp")

		return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	// save user
	_, err = s.userSvc.CreateUser(ctx, &dto.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     enum.RoleUser,
	})
	if err != nil {
		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
			"req":   loggableReq,
		}, "[AuthService][RegisterUser] failed to create user")

		return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	log.Info(map[string]interface{}{
		"email": req.Email,
	}, "[AuthService][RegisterUser] user registered")

	// login user
	return s.LoginUser(ctx, dto.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
}

func (s *authService) LoginUser(ctx context.Context,
	req dto.LoginUserRequest) (dto.LoginResponse, error) {

	var resp dto.LoginResponse

	// get user by email
	user, err := s.userSvc.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, errorpkg.ErrNotFound) {
			return resp, errorpkg.ErrNotFound.WithMessage("User not found. Please register first.")
		}

		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
			"email": req.Email,
		}, "[AuthService][LoginUser] failed to get user by email")

		return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	// check password
	ok := s.bcrypt.Compare(req.Password, user.PasswordHash)
	if !ok {
		return resp, errorpkg.ErrCredentialsNotMatch
	}

	// Create channels for token generation results
	type tokenResult struct {
		token string
		err   error
	}
	accessTokenCh := make(chan tokenResult)
	refreshTokenCh := make(chan tokenResult)

	// Generate access token
	go func() {
		token, err := s.jwt.Create(user.ID, user.Role)
		accessTokenCh <- tokenResult{token: token, err: err}
	}()

	// Generate and store refresh token
	go func() {
		refreshToken := randgen.RandomString(32)
		err := s.repo.CreateSession(ctx, &entity.Session{
			Token:     refreshToken,
			UserID:    user.ID,
			ExpiresAt: time.Now().Add(env.GetEnv().JwtRefreshExpireDuration),
		})
		refreshTokenCh <- tokenResult{token: refreshToken, err: err}
	}()

	var accessResult, refreshResult tokenResult
	resultsReceived := 0

	// Wait for both operations to complete in any order
	for resultsReceived < 2 {
		select {
		case accessResult = <-accessTokenCh:
			if accessResult.err != nil {
				traceID := log.ErrorWithTraceID(map[string]interface{}{
					"error": accessResult.err.Error(),
					"email": req.Email,
				}, "[AuthService][LoginUser] failed to generate access token")
				return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
			}
			resultsReceived++

		case refreshResult = <-refreshTokenCh:
			if refreshResult.err != nil {
				traceID := log.ErrorWithTraceID(map[string]interface{}{
					"error": refreshResult.err.Error(),
					"email": req.Email,
				}, "[AuthService][LoginUser] failed to store session")
				return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
			}
			resultsReceived++
		}
	}

	userResp := dto.UserResponse{}
	userResp.PopulateFromEntity(user)
	resp = dto.LoginResponse{
		AccessToken:  accessResult.token,
		RefreshToken: refreshResult.token,
		User:         &userResp,
	}

	log.Info(map[string]interface{}{
		"email": req.Email,
	}, "[AuthService][LoginUser] user logged in")

	return resp, nil
}

func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (dto.LoginResponse, error) {
	var resp dto.LoginResponse

	session, err := s.repo.GetSessionByToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return resp, errorpkg.ErrInvalidRefreshToken
		}

		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error": err.Error(),
		}, "[AuthService][RefreshToken] failed to get session by token")

		return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	if session.ExpiresAt.Before(time.Now()) {
		return resp, errorpkg.ErrInvalidRefreshToken
	}

	// get user by session
	user, err := s.userSvc.GetUserByID(ctx, session.UserID)
	if err != nil {
		if errors.Is(err, errorpkg.ErrNotFound) {
			return resp, errorpkg.ErrNotFound.WithMessage("User not found. Please register first.")
		}

		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error":  err.Error(),
			"userID": session.UserID,
		}, "[AuthService][RefreshToken] failed to get user by ID")

		return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	accessToken, err := s.jwt.Create(user.ID, user.Role)
	if err != nil {
		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error":   err.Error(),
			"user_id": user.ID,
		}, "[AuthService][RefreshToken] failed to generate access token")

		return resp, errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	userResp := dto.UserResponse{}
	userResp.PopulateFromEntity(user)
	resp = dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         &userResp,
	}

	log.Info(map[string]interface{}{
		"user.id":    user.ID,
		"user.email": user.Email,
	}, "[AuthService][RefreshToken] token refreshed")

	return resp, nil
}

func (s *authService) Logout(ctx context.Context) error {
	userID := ctx.Value("user.id").(uuid.UUID)

	err := s.repo.DeleteSession(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errorpkg.ErrInvalidBearerToken
		}

		traceID := log.ErrorWithTraceID(map[string]interface{}{
			"error":  err.Error(),
			"userID": userID,
		}, "[AuthService][Logout] failed to delete session")

		return errorpkg.ErrInternalServer.WithTraceID(traceID)
	}

	log.Info(map[string]interface{}{
		"user.id": userID,
	}, "[AuthService][Logout] user logged out")

	return nil
}
