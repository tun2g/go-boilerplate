package post


import (
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/apis/model"
	"fist-app/src/apis/dto/post"
	pageDto "fist-app/src/shared/dto"
)

type PostService interface{
	CreateNewPost(ctx *httpContext.CustomContext, userId string, dto post.CreatePostReqDto) (*model.Post, error)
	GetPostsByUserId(ctx *httpContext.CustomContext, userId string, dto *pageDto.PageOptionsDto) (*pageDto.PageDto, error)
}