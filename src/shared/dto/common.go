package dto

import (
	"fist-app/src/shared/jwt"
)

type CurrentUser struct {
	FullName string `json:"fullName,omitempty"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}

func NewCurrentUser(payload *jwt.JwtPayload) *CurrentUser {
	return &CurrentUser{
		FullName: payload.FullName,
		Email: payload.Email,
		ID: payload.UserId,
	}
}
