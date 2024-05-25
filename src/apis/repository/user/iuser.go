package repository

import "fist-app/src/apis/model"

type UserRepository interface {
	FindUserByEmail(email string) (model.User, error)
	FindUserByID(ID int) model.User
	StoreUser(user model.User) (model.User, error)
}
