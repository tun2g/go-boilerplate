package repository

import (
	"errors"
	"fist-app/src/apis/model"
	"fist-app/src/shared/dto"
	pageDto "fist-app/src/shared/dto"
	"fmt"

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

func (repo usersRepository) GetAll(dto *pageDto.PageOptionsDto) (*[]model.User, error){
	var users []model.User
	query :=  repo.storage.
		Offset(*dto.Offset).
		Limit(*dto.Limit)
	
	if(dto.Order != nil){
		orderField := "createdAt"
		if(dto.OrderField != nil){
			dto.OrderField = &orderField
		}
		query.Order(fmt.Sprintf("%s %s",*dto.OrderField, *dto.Order))
	}

	query.Find(&users)

	err := query.Error
	
	if(err!= nil){
		return nil, err
	}
	return &users, nil
}

func (repo usersRepository) CountByPageDto(dto *dto.PageOptionsDto) (int, error){
	var count int64
	query := repo.storage.Model(&model.User{})

	err := query.Limit(*dto.Limit).Offset(*dto.Offset).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil	
}