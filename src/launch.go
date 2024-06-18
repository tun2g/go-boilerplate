package src

import (
	authController "fist-app/src/apis/controller/auth"
	authService "fist-app/src/apis/service/auth"
	userRepository "fist-app/src/apis/repository/user"

	postController "fist-app/src/apis/controller/post"
	postService "fist-app/src/apis/service/post"
	postRepository "fist-app/src/apis/repository/post"

	"fist-app/src/shared/jwt"
	"fist-app/src/shared/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) launchingServer(route *gin.Engine) {
	var jwtAccessTokenManager = jwt.NewJWTManager(
		server.config.JwtAccessTokenSecret,
		time.Duration(server.config.JwtAccessTokenExpirationTime),
	)

	var jwtRefreshTokenManager = jwt.NewJWTManager(
		server.config.JwtRefreshTokenSecret,
		time.Duration(server.config.JwtRefreshTokenExpirationTime),
	)

	var bcrypt = utils.NewBcryptEncoder(bcrypt.DefaultCost)

	// Auth module
	var _authService = authService.NewAuthService(
		userRepository.NewUsersRepository(server.db),
		&jwtAccessTokenManager,
		&jwtRefreshTokenManager,
		&bcrypt,
	)
	var _authController = authController.NewAuthController(server.ctx, _authService)
	authRoutes := route.Group("/auth")
	_authController.InitRoute(authRoutes, &jwtAccessTokenManager, &jwtRefreshTokenManager)


	// Post Module
	var _postService = postService.NewPostService(
		postRepository.NewPostRepository(server.db),
	)

	var _postController = postController.NewPostController(server.ctx, _postService)
	postRoutes :=route.Group("/posts")
	_postController.InitRoute(postRoutes, &jwtAccessTokenManager)
}
