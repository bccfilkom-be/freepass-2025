package model

import "time"

type User struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	FullName    string    `json:"full_name"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	Affiliation string    `json:"affiliation"`
	IsVerified  bool      `json:"is_verified"`
	VerifiedAt  time.Time `json:"verified_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserCreate struct {
	Email       string `json:"email" validate:"required,email" example:"rchronicler@gmail.com"`
	Password    string `json:"password" validate:"required,min=8" example:"password123"`
	FullName    string `json:"full_name" validate:"required" example:"John Doe"`
	Affiliation string `json:"affiliation" validate:"required" example:"RAION"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" example:"rchronicler@gmail.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type UserLoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type UserProfileUpdate struct {
	FullName    string `json:"full_name" validate:"required" example:"Dimas Andrian"`
	Affiliation string `json:"affiliation" validate:"required" example:"BCC"`
}
