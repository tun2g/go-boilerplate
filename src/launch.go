package src

import (
	authController "fist-app/src/apis/controller/auth"
	authService "fist-app/src/apis/service/auth"
	repository "fist-app/src/apis/repository/user"
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
	var authService = authService.NewAuthService(
		repository.NewUsersRepository(server.db),
		&jwtAccessTokenManager,
		&jwtRefreshTokenManager,
		&bcrypt,
	)
	var authController = authController.NewAuthController(server.ctx, authService)
	authRoutes := route.Group("/auth")
	authController.InitRoute(authRoutes, &jwtAccessTokenManager, &jwtRefreshTokenManager)

}
