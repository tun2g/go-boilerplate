package post

import (
	"fist-app/src/apis/model"
	postDto "fist-app/src/apis/dto/post"
	pageDto "fist-app/src/shared/dto"
)

type PostRepository interface {
	StorePost(post *model.Post) (*model.Post, error)
	GetPostsByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto) (*[]model.Post, error)
	CountByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto) (int, error)
	SoftDeletePost(userId string, postId string) (error)
	UpdatePost(userId string, postId string, updatedPost *postDto.UpdatePostReqDto) (*model.Post, error)
	GetPost(userId string, postId string) (*model.Post, error)
}