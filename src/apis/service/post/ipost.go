package post

import (
	"fist-app/src/apis/dto/post"
	postDto "fist-app/src/apis/dto/post"
	"fist-app/src/apis/model"
	pageDto "fist-app/src/shared/dto"
	httpContext "fist-app/src/shared/http-context"
)

type PostService interface {
	CreateNewPost(ctx *httpContext.CustomContext, userId string, dto post.CreatePostReqDto) (*model.Post, error)
	GetPostsByUserId(ctx *httpContext.CustomContext, userId string, dto *pageDto.PageOptionsDto) (*pageDto.PageDto, error)
	SoftDeletePost(userId string, postId string) error
	UpdatePost(userId string, postId string, updatedPost *postDto.UpdatePostReqDto) (*model.Post, error)
	GetPost(userId string, postId string) (*model.Post, error)
}
