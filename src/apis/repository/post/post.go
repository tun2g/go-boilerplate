package post

import (
	"fist-app/src/apis/model"

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
