package auth

import (
	"context"
	"net/http"

	dto "fist-app/src/apis/dto/auth"
	pageDto "fist-app/src/shared/dto"
	authService "fist-app/src/apis/service/auth"
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/apis/model"
	"fist-app/src/shared/exception"
)

type AuthController struct {
	ctx         context.Context
	authService authService.AuthService
}

func NewAuthController(ctx context.Context, authService authService.AuthService) *AuthController {
	return &AuthController{
		ctx:         ctx,
		authService: authService,
	}
}


func (handler *AuthController) Login(ctx *httpContext.CustomContext) {
	var reqDto dto.LoginReqDto
	var err error
	var user *model.User
	var tokens *dto.TokenResDto

	if err := ctx.ShouldBindJSON(&reqDto); err != nil {
		ctx.Error(exception.NewUnprocessableEntityException(ctx.GetRequestId(), err))
		return
	}

	user, tokens, err = handler.authService.Login(reqDto, ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	userRes := dto.UserResDto{
		ID:       user.Id,
		Email:    user.Email,
		FullName: user.FullName,
	}

	authRes := dto.AuthResDto{
		User:   userRes,
		Tokens: *tokens,
	}

	ctx.JSON(http.StatusOK, authRes)
}


func (handler *AuthController) Register(ctx *httpContext.CustomContext) {
	var reqDto dto.RegisterReqDto
	var err error
	var user *model.User
	var tokens *dto.TokenResDto

	if err := ctx.ShouldBindJSON(&reqDto); err != nil {
		ctx.Error(exception.NewUnprocessableEntityException(ctx.GetRequestId(), err))
		return
	}

	user, tokens, err = handler.authService.Register(reqDto, ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	userRes := dto.UserResDto{
		ID:       user.Id,
		Email:    user.Email,
		FullName: user.FullName,
	}

	authRes := dto.AuthResDto{
		User:   userRes,
		Tokens: *tokens,
	}

	ctx.JSON(http.StatusCreated, authRes)
}


func (handler *AuthController) GetMe(ctx *httpContext.CustomContext) {
	user := handler.authService.GetMe(ctx)

	if user == nil {
		ctx.Error(exception.NewBadRequestException(
			ctx.GetRequestId(),
			[]exception.ErrorDetail{{}},
		))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (handler *AuthController) RefreshToken(ctx *httpContext.CustomContext) {
	tokens, err := handler.authService.RefreshToken(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *tokens)
}


func (handler *AuthController) GetUsers(ctx *httpContext.CustomContext){
	var queryDto pageDto.PageOptionsDto
	if err := ctx.ShouldBindQuery(&queryDto); err != nil {
		ctx.Error(exception.NewUnprocessableEntityException(ctx.GetRequestId(), err))
		return
	}

	queryDto = *pageDto.NewPageOptionsDto(&queryDto);

	data, err := handler.authService.GetUsers(ctx, &queryDto);

	if(err!=nil){
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}