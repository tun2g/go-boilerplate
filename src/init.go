package src

import (
	authController "fist-app/src/apis/controller/auth"
	repository "fist-app/src/apis/repository/user"
	authService "fist-app/src/apis/service/auth"
)

func (server * Server) InitServer() map[string]interface{}{
	var authService = authService.NewAuthService(repository.NewUsersRepository(server.db))
	
	var authController = authController.NewAuthController(server.ctx ,authService)

	return map[string]interface{}{
		"auth": authController,
	}
}