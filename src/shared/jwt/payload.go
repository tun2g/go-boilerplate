package jwt

import (
	"fist-app/src/apis/model"
	"time"

	"github.com/google/uuid"
)

type JwtPayload struct {
	UserId    string    `json:"userId"`
	Email     string    `json:"email"`
	TokenId   string    `json:"tokenId"`
	FullName  string    `json:"fullName"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredTime"`
}

func NewJwtPayload(user *model.User, expirationTime time.Duration) *JwtPayload {
	tokenId := uuid.New().String()
	return &JwtPayload{
		TokenId:   tokenId,
		UserId:    user.ID,
		Email:     user.Email,
		Role:      user.Role,
		FullName:  user.FullName,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(expirationTime * 1000000000),
	}
}

func (payload *JwtPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
