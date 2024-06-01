package auth

import (
	"fist-app/src/apis/dto/auth"
	"fist-app/src/apis/model"
	repository "fist-app/src/apis/repository/user"
	"fist-app/src/shared/dto"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/shared/jwt"
	"fist-app/src/shared/utils"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepository         repository.UserRepository
	jwtAccessTokenManager  *jwt.JWTManager
	jwtRefreshTokenManager *jwt.JWTManager
	bcrypt                 *utils.BcryptEncoder
}

func NewAuthService(
	userRepository repository.UserRepository,
	jwtAccessTokenManager *jwt.JWTManager,
	jwtRefreshTokenManager *jwt.JWTManager,
	bcrypt *utils.BcryptEncoder,
) *authService {
	return &authService{
		userRepository:         userRepository,
		jwtAccessTokenManager:  jwtAccessTokenManager,
		jwtRefreshTokenManager: jwtRefreshTokenManager,
		bcrypt:                 bcrypt,
	}
}

func (srv *authService) Register(req auth.RegisterReqDto, ctx *httpContext.CustomContext) (*model.User, *auth.TokenResDto, error) {
	var err error

	user, err := srv.userRepository.FindUserByEmail(req.Email)

	if user != nil {
		err = exception.NewBadRequestException(
			ctx.GetRequestId(),
			[]exception.ErrorDetail{{
				Issue:   "Email is already in use",
				Field:   "email",
				IssueId: "exists_email",
			}},
		)
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}

	hashedPassword, err := srv.bcrypt.Encrypt(req.Password)

	user, err = srv.userRepository.StoreUser(model.User{
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
	})

	if err != nil {
		return nil, nil, err
	}

	accessToken, _, err := srv.jwtAccessTokenManager.CreateToken(user)
	refreshToken, _, err := srv.jwtRefreshTokenManager.CreateToken(user)
	tokens := auth.TokenResDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return user, &tokens, nil
}

func (srv *authService) Login(req auth.LoginReqDto, ctx *httpContext.CustomContext) (*model.User, *auth.TokenResDto, error) {
	var err error

	user, err := srv.userRepository.FindUserByEmail(req.Email)
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		err = exception.NewBadRequestException(
			ctx.GetRequestId(),
			[]exception.ErrorDetail{{
				Issue:   "Email or password is invalid",
			}},
		)
		return nil, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, nil, err
	}

	accessToken, _, err := srv.jwtAccessTokenManager.CreateToken(user)
	refreshToken, _, err := srv.jwtRefreshTokenManager.CreateToken(user)
	tokens := auth.TokenResDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return user, &tokens, nil
}

func (srv *authService) GetMe(ctx *httpContext.CustomContext) *dto.CurrentUser {

	user := ctx.GetUser()

	return user
}

func (srv *authService) RefreshToken(ctx *httpContext.CustomContext) (*auth.TokenResDto, error){
	user := ctx.GetUser()

	if user == nil {
		err := exception.NewUnauthorizedException(ctx.GetRequestId())
		return nil, err
	}

	_user, err := srv.userRepository.FindUserByEmail(user.Email)

	if user == nil {
		err := exception.NewUnauthorizedException(ctx.GetRequestId())
		return nil, err
	}

	if err != nil {
		err := exception.NewUnauthorizedException(ctx.GetRequestId())
		return nil, err
	}

	accessToken, _, err := srv.jwtAccessTokenManager.CreateToken(_user)
	refreshToken, _, err := srv.jwtRefreshTokenManager.CreateToken(_user)
	tokens := auth.TokenResDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &tokens, nil
}