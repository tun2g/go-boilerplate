package auth

import (
	"context"
	"net/http"

	dto "fist-app/src/apis/dto/auth"
	"fist-app/src/apis/model"
	"fist-app/src/apis/service/auth"
	"fist-app/src/shared/exception"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	ctx         context.Context
	authService auth.AuthService
}

func NewAuthController(ctx context.Context, authService auth.AuthService) *AuthController {
	return &AuthController{
		ctx:         ctx,
		authService: authService,
	}
}

func (handler *AuthController) Login(context * gin.Context){
		var reqDto dto.LoginReqDto
		var err error
		var user model.User
		err = context.ShouldBind(&reqDto)

		if err != nil {
			exception.NewUnprocessableEntityException(
				context.GetString("requestId"),
				[]exception.ErrorDetail{{
					Issue:   "",
					Message: "a",
				}},
			)
			return
		}

		user, err = handler.authService.Login(reqDto)

		if err != nil {
			exception.NewBadRequestException(
				context.GetString("requestId"),
				[]exception.ErrorDetail{{
					Issue:   "error",
					Message: "error",
				}},
			)
			return
		}
		context.JSON(http.StatusCreated, user)
}
