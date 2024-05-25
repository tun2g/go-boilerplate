package src

import (
	"github.com/gin-gonic/gin"
	"fist-app/src/apis/controller/auth"
)


func (server *Server) setupRoutes(route *gin.Engine){
	controller := server.InitServer()

	authRoutes := route.Group("auth")
	authController := controller["auth"].(*auth.AuthController)
	authController.InitRoute(authRoutes)	
}