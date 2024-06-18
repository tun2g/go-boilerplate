package post

import (
	"fist-app/src/apis/model"
	pageDto "fist-app/src/shared/dto"
)

type PostRepository interface {
	StorePost(post *model.Post) (*model.Post, error)
	GetPostsByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto) (*[]model.Post, error)
	CountByUserIdAndPageDto(userId string, dto *pageDto.PageOptionsDto) (int, error)
}