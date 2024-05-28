package auth

import (
	authConstants "fist-app/src/shared/auth/constants"
	"fist-app/src/shared/dto"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/shared/jwt"
	"strings"
)

func AuthMiddleware(jwtManager *jwt.JWTManager) func(ctx *httpContext.CustomContext) {
	return func(ctx *httpContext.CustomContext) {
		authorizationHeader := ctx.GetHeader(authConstants.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			exception.ThrowUnauthorizedException(ctx)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			exception.ThrowUnauthorizedException(ctx)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authConstants.AuthorizationBearerKey {
			exception.ThrowUnauthorizedException(ctx)
			return
		}

		accessToken := fields[1]
		payload, err := jwtManager.VerifyToken(accessToken)
		if err != nil {
			exception.ThrowUnauthorizedException(ctx)
			return
		}

		currentUser := dto.NewCurrentUser(payload)

		ctx.SetUser(currentUser)
		ctx.Next()
	}
}