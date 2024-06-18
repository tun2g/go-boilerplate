package repository

import (
	"fist-app/src/apis/model"
	pageDto "fist-app/src/shared/dto"
)

type UserRepository interface {
	FindUserByEmail(email string) (*model.User, error)
	FindUserByID(id string) (*model.User, error)
	StoreUser(user model.User) (*model.User, error)
	GetAll(dto *pageDto.PageOptionsDto) (*[]model.User, error)
	CountByPageDto(dto *pageDto.PageOptionsDto) (int, error)
}
