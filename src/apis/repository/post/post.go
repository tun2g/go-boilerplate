package post

import (
	"errors"
	postDto "fist-app/src/apis/dto/post"
	"fist-app/src/apis/model"
	pageDto "fist-app/src/shared/dto"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type postRepository struct {
	storage *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{storage: db}
}

func (repo *postRepository) StorePost(post *model.Post) (*model.Post, error) {
	err := repo.storage.Create(&post).Error
	return post, err
}

func (repo *postRepository) GetPostsByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto) (*[]model.Post, error) {
	var posts []model.Post
	query := repo.storage.
		Where("`user_id` = ?", userId).
		Offset(*dto.Offset).
		Limit(*dto.Limit)

	if dto.Order != nil {
		orderField := "createdAt"
		if dto.OrderField != nil {
			dto.OrderField = &orderField
		}
		query.Order(fmt.Sprintf("%s %s", *dto.OrderField, *dto.Order))
	}

	query.Find(&posts)

	err := query.Error

	if err != nil {
		return nil, err
	}
	return &posts, nil
}

func (repo *postRepository) CountByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto) (int, error) {
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

// SoftDeletePost performs a soft delete on a post by updating the DeletedAt field
func (repo *postRepository) SoftDeletePost(userId string, postId string) error {
	result := repo.storage.Model(&model.Post{}).Where("user_id = ? AND id = ?", userId, postId).Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found or already deleted")
	}
	return nil
}

// UpdatePost updates a post's information
func (repo *postRepository) UpdatePost(userId string, postId string, updatedPost *postDto.UpdatePostReqDto) (*model.Post, error) {
	var post model.Post
	if err := repo.storage.Where("user_id = ? AND id = ?", userId, postId).First(&post).Error; err != nil {
		return nil, err
	}

	// Update the post fields
	post.Title = updatedPost.Title
	post.Description = updatedPost.Description

	if err := repo.storage.Save(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// GetPost retrieves a post by userId and postId
func (repo *postRepository) GetPost(userId string, postId string) (*model.Post, error) {
	var post model.Post
	if err := repo.storage.Where("user_id = ? AND id = ? AND deleted_at IS NULL", userId, postId).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
