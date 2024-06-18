package post

import (
	postDto "fist-app/src/apis/dto/post"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
)

func initPostDto(_ postDto.PostResDto)    {}
func initException(_ exception.HttpError) {}

// @Summary Create a post
// @Description Create a post
// @Tags Post
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   createPostReq  body  postDto.CreatePostReqDto  true "Create Post Request"
// @Success 201 {object} postDto.PostResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Failure 401 {object} exception.HttpError
// @Failure 500 {object} exception.HttpError
// @Router /posts [post]
func createPost(ctx *httpContext.CustomContext) {}
