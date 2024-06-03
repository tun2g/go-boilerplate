package jwt

import (
	"fist-app/src/apis/model"
)

type Manager interface {
	CreateToken(user *model.User) (string, *JwtPayload, error)

	VerifyToken(token string) (*JwtPayload, error)
}
