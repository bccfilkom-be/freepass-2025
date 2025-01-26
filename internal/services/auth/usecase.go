package auth

import "jevvonn/bcc-be-freepass-2025/internal/models/dto"

type AuthUsecase interface {
	SignUp(*dto.SignUpRequest) error
}
