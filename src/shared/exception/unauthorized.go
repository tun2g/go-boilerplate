package exception

import (
	httpContext "fist-app/src/shared/http-context"
	"net/http"
)

type UnauthorizedException struct {
	HttpError
}

func NewUnauthorizedException(requestId string) *UnauthorizedException {
	return &UnauthorizedException{
		HttpError: HttpError{
			Message: "Unauthorized",
			RequestId: requestId,
			Details: []ErrorDetail{{}},
		},
	}
}

func ThrowUnauthorizedException(ctx *httpContext.CustomContext){
	ctx.AbortWithStatusJSON(
		http.StatusUnauthorized,
		HttpError{
			Message: "Unauthorized",
			RequestId: ctx.GetRequestId(),
			Details: []ErrorDetail{{}},
		},
	)
}