package dto

import (
	"fist-app/src/shared/jwt"
)

type CurrentUser struct {
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email"`
	ID       string `json:"id"`
	Role     string `json:"role"`
}

func NewCurrentUser(payload *jwt.JwtPayload) *CurrentUser {
	if payload != nil {
		return &CurrentUser{
			FullName: payload.FullName,
			Email:    payload.Email,
			ID:       payload.UserId,
			Role:     payload.Role,
		}
	}
	return nil
}
