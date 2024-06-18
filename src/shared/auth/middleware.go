package auth

import (
	authConstants "fist-app/src/shared/auth/constants"
	"fist-app/src/shared/dto"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/shared/jwt"
	"fist-app/src/shared/utils"
	"strings"
	"fist-app/src/config"
)

func IsPublicRouteMiddleware() func(ctx *httpContext.CustomContext){
	return func(ctx *httpContext.CustomContext){
		ctx.SetIsPublicRoute();
	}
}

func TokenAuthMiddleware(jwtManager *jwt.JWTManager) func(ctx *httpContext.CustomContext) {
	return func(ctx *httpContext.CustomContext) {

		isPublic := ctx.GetIsPublicRoute();
		var accessToken string
		authorizationHeader := ctx.GetHeader(authConstants.AuthorizationHeaderKey)
		
		if(config.AppConfiguration.GoEnv == "production"){
			if len(authorizationHeader) == 0 && isPublic == false {
				exception.ThrowUnauthorizedException(ctx)
				return
			}
	
			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 && isPublic == false{
				exception.ThrowUnauthorizedException(ctx)
				return
			}
	
			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authConstants.AuthorizationBearerKey && isPublic == false {
				exception.ThrowUnauthorizedException(ctx)
				return
			}
			accessToken = fields[1]

		}else {
			// PURPOSE: enable swagger authorization 
			accessToken = strings.Fields(authorizationHeader)[0]
		}


		payload, err := jwtManager.VerifyToken(accessToken)
		if err != nil && isPublic == false{
			exception.ThrowUnauthorizedException(ctx)
			return
		}

		currentUser := dto.NewCurrentUser(payload)

		ctx.SetUser(currentUser)
		ctx.Next()
	}
}

func RoleAuthMiddleware(roles []string) func(ctx *httpContext.CustomContext){
	return func(ctx *httpContext.CustomContext) {
		user := ctx.GetUser()
		if(user == nil){
			exception.ThrowForbiddenException(ctx)
			return
		}

		if(!utils.IsContains(roles, user.Role)){
			exception.ThrowForbiddenException(ctx)
			return
		}
		ctx.Next()
	}
}