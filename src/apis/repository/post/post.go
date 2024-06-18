package post

import (
	"fist-app/src/apis/model"
	pageDto "fist-app/src/shared/dto"
	"fmt"

	"gorm.io/gorm"
)

type postRepository struct {
	storage *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return postRepository{storage: db}
}


func (repo postRepository) StorePost(post *model.Post) (*model.Post, error){
	err := repo.storage.Create(&post).Error
	return post, err;
}

func (repo postRepository) GetPostsByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto)(*[]model.Post, error){
	var posts []model.Post
	query :=  repo.storage.
		Where("`user_id` = ?", userId).
		Offset(*dto.Offset).
		Limit(*dto.Limit)
	
	if(dto.Order != nil){
		orderField := "createdAt"
		if(dto.OrderField != nil){
			dto.OrderField = &orderField
		}
		query.Order(fmt.Sprintf("%s %s",*dto.OrderField, *dto.Order))
	}

	query.Find(&posts)

	err := query.Error
	
	if(err!= nil){
		return nil, err
	}
	return &posts, nil
}

func (repo postRepository) CountByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto)(int, error){
	var count int64
	query := repo.storage.Model(&model.Post{})

	err := query.Where("user_id = ?", userId).
		Limit(*dto.Limit).
		Offset(*dto.Offset).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}
