package auth

import (
	"fist-app/src/apis/dto/auth"
	"fist-app/src/apis/model"
	"fist-app/src/shared/dto"
	pageDto "fist-app/src/shared/dto"
	httpContext "fist-app/src/shared/http-context"
)

type AuthService interface {
	Login(dto auth.LoginReqDto, ctx *httpContext.CustomContext) (*model.User, *auth.TokenResDto, error)
	Register(dto auth.RegisterReqDto, ctx *httpContext.CustomContext) (*model.User, *auth.TokenResDto, error)
	GetMe(ctx *httpContext.CustomContext) *dto.CurrentUser
	RefreshToken(ctx *httpContext.CustomContext) (*auth.TokenResDto, error)
	GetUsers(ctx *httpContext.CustomContext, dto *pageDto.PageOptionsDto) (*pageDto.PageDto, error)
}
