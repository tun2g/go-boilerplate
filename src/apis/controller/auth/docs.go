package auth

import (
	authDto "fist-app/src/apis/dto/auth"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
)

func initAuthDto(_ authDto.AuthResDto)    {}
func initException(_ exception.HttpError) {}

// @Summary Login
// @Description User login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   loginReq  body  authDto.LoginReqDto  true  "Login request"
// @Success 200 {object} authDto.AuthResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Router /auth/sign-in [post]
func login(ctx *httpContext.CustomContext) {}

// @Summary Register
// @Description User Register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   registerReq  body  authDto.RegisterReqDto  true  "Register request"
// @Success 201 {object} authDto.AuthResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Router /auth/sign-up [post]
func register(ctx *httpContext.CustomContext) {}
