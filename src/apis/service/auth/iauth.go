package auth

import (
	"fist-app/src/apis/dto/auth"
	"fist-app/src/apis/model"
)

type AuthService interface {
	Login(dto auth.LoginReqDto) (model.User ,error)
}