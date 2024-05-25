package auth

import (
	"fist-app/src/apis/dto/auth"
	"fist-app/src/apis/model"
	"fist-app/src/apis/repository/user"
)

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *authService {
	return &authService{
		userRepository: userRepository,
	}
}

func (srv *authService) Login(req auth.LoginReqDto) (model.User, error) {
	var err error

	user, err := srv.userRepository.StoreUser(model.User{
		Email:    req.Email,
		Password: req.Password,
		FullName: "",
	})

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
