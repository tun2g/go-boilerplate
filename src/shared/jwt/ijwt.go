package jwt

import "time"

type Manager interface {
	CreateToken(username string, duration time.Duration) (string, *JwtPayload, error)

	VerifyToken(token string) (*JwtPayload, error)
}
