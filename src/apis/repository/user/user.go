package repository

import (
	"errors"
	"fist-app/src/apis/model"

	"gorm.io/gorm"
)

type usersRepository struct {
	storage *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UserRepository {
	return usersRepository{storage: db}
}

func (repo usersRepository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.storage.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if(err != nil){
		return nil, err
	}

	return &user, nil
}

func (repo usersRepository) FindUserByID(id int) (*model.User, error) {
	var user model.User
	err := repo.storage.Where("id = ?", id).First(&user).Error
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, nil
}

func (repo usersRepository) StoreUser(user model.User) (*model.User, error) {
	err := repo.storage.Create(&user).Error
	return &user, err
}
