package repository

import (
	"fist-app/src/apis/model"

	"github.com/jinzhu/gorm"
)

type usersRepository struct {
	storage *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UserRepository {
	return usersRepository{storage: db}
}

func (repo usersRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	err := repo.storage.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return handleErr(err)
	}

	return user, nil
}

func (repo usersRepository) FindUserByID(id int) model.User {
	var user model.User
	repo.storage.Where("id = ?", id).First(&user)

	return user
}

func (repo usersRepository) StoreUser(user model.User) (model.User, error) {
	err := repo.storage.Create(&user).Error
	return user, err
}

func handleErr(err error) (model.User, error) {
	if gorm.IsRecordNotFoundError(err) {
		return model.User{}, nil
	}

	return model.User{}, err
}
