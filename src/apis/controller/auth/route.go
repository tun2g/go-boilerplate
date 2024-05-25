package auth

import "github.com/gin-gonic/gin"

func (authController *AuthController) InitRoute(routes *gin.RouterGroup) {
	routes.POST("/login", authController.Login)
}
