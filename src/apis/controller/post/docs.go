package post

import (
	postDto "fist-app/src/apis/dto/post"
	pageDto "fist-app/src/shared/dto"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
)

func initPostDto(_ postDto.PostResDto)    {}
func initException(_ exception.HttpError) {}
func initPageDto(_ pageDto.PageDto)       {}

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

// @Summary Get posts by user ID with pagination and ordering
// @Description Retrieves posts for a specific user based on provided query parameters.
// @Tags Post
// @Param pageOptions query pageDto.PageOptionsDto true "Pagination and ordering options"
// @Produce json
// @Security BearerAuth
// @Success 200 {array} postDto.PostResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Failure 401 {object} exception.HttpError
// @Failure 500 {object} exception.HttpError
// @Router /posts [get]
func getPosts(ctx *httpContext.CustomContext) {}
