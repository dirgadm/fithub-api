package dto

import "time"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Name     string `json:"name" validate:"required,gte=2"`
	Phone    string `json:"phone" validate:"required,gte=12"`
}

type AuthResponse struct {
	Token     string       `json:"token"`
	ExpiredAt time.Time    `json:"expired_at"`
	User      UserResponse `json:"user"`
}
