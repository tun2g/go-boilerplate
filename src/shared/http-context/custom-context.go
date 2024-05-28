package httpContext

import (
	"fist-app/src/shared/dto"
	authConstants "fist-app/src/shared/auth/constants"

	"github.com/gin-gonic/gin"
	httpContextConstant "fist-app/src/shared/http-context/constants"
)

type CustomContext struct {
	*gin.Context
}

func (ctx *CustomContext) GetUser() *dto.CurrentUser {
	user, exists := ctx.Get(authConstants.AuthUser)
	if !exists {
		return nil
	}
	if userStruct, ok := user.(*dto.CurrentUser); ok {
		return userStruct
	}
	return nil
}

func (ctx *CustomContext) SetUser(user *dto.CurrentUser){
	ctx.Set(authConstants.AuthUser, user)
}

func (ctx *CustomContext) SetRequestId(requestId string){
	ctx.Set(httpContextConstant.RequestId, requestId)
}

func (ctx * CustomContext) GetRequestId() string {
	requestId:= ctx.GetString(httpContextConstant.RequestId);
	return requestId
}

