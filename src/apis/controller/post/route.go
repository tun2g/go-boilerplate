package post

import (
	"fist-app/src/shared/auth"
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/shared/jwt"

	"github.com/gin-gonic/gin"
)

func (postController *PostController) InitRoute(
	routes *gin.RouterGroup,
	jwtAccessTokenManager *jwt.JWTManager,
) {
	routes.POST(
		"",
		httpContext.CustomContextHandler(auth.TokenAuthMiddleware(jwtAccessTokenManager)),
		httpContext.CustomContextHandler(postController.CreateNewPost),
	)

	routes.GET(
		"",
		httpContext.CustomContextHandler(auth.TokenAuthMiddleware(jwtAccessTokenManager)),
		httpContext.CustomContextHandler(postController.GetPostsByUser),
	)
}
