package auth

import (
	"fist-app/src/shared/auth"
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/shared/jwt"

	"github.com/gin-gonic/gin"
)

func (authController *AuthController) InitRoute(
	routes *gin.RouterGroup,
	jwtAccessTokenManager *jwt.JWTManager,
	jwtRefreshTokenManager *jwt.JWTManager,
) {
	routes.POST(
		"/sign-in", 
		httpContext.CustomContextHandler(authController.Login),
	)

	routes.POST(
		"/sign-up",
		httpContext.CustomContextHandler(authController.Register),
	)

	routes.GET(
		"/me",
		httpContext.CustomContextHandler(auth.TokenAuthMiddleware(jwtAccessTokenManager)),
		httpContext.CustomContextHandler(authController.GetMe),
	)

	routes.GET(
		"/refresh-token",
		httpContext.CustomContextHandler(auth.TokenAuthMiddleware(jwtRefreshTokenManager)),
		httpContext.CustomContextHandler(authController.RefreshToken),
	)

	routes.GET(
		"/all",
		httpContext.CustomContextHandler(authController.GetUsers),
	)
}
