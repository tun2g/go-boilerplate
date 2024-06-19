package post

import (
	"context"
	postDto "fist-app/src/apis/dto/post"
	"fist-app/src/apis/model"
	postService "fist-app/src/apis/service/post"
	"fist-app/src/shared/dto"
	pageDto "fist-app/src/shared/dto"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
	"net/http"
)

type PostController struct {
	ctx         context.Context
	postService postService.PostService
}

func NewPostController(ctx context.Context, postService postService.PostService) *PostController {
	return &PostController{
		ctx:         ctx,
		postService: postService,
	}
}

func (handler *PostController) CreateNewPost(ctx *httpContext.CustomContext) {
	var reqDto postDto.CreatePostReqDto
	var err error
	var post *model.Post

	if err := ctx.ShouldBindJSON(&reqDto); err != nil {
		ctx.Error(exception.NewUnprocessableEntityException(ctx.GetRequestId(), err))
		return
	}

	user := ctx.GetUser()
	post, err = handler.postService.CreateNewPost(ctx, user.Id, reqDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	postRes := postDto.PostResDto{
		AuditableResDto: dto.AuditableResDto{
			Id:        post.Id,
			CreatedAt: post.CreatedAt,
			DeletedAt: post.DeletedAt.Time,
			UpdatedAt: post.UpdatedAt,
		},
		Title:       post.Title,
		Description: post.Description,
		UserId:      post.UserId,
	}

	ctx.JSON(http.StatusCreated, postRes)
}

func (handler *PostController) GetPostsByUser(ctx *httpContext.CustomContext) {
	var queryDto pageDto.PageOptionsDto
	if err := ctx.ShouldBindQuery(&queryDto); err != nil {
		ctx.Error(exception.NewUnprocessableEntityException(ctx.GetRequestId(), err))
		return
	}
	queryDto = *pageDto.NewPageOptionsDto(&queryDto)
	user := ctx.GetUser()

	data, err := handler.postService.GetPostsByUserId(ctx, user.Id, &queryDto)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (handler *PostController) GetPostById(ctx *httpContext.CustomContext) {
	postId := ctx.Param("id")
	user := ctx.GetUser()

	post, err := handler.postService.GetPost(user.Id, postId)

	if err != nil {
		ctx.Error(err)
		return
	}

	postRes := postDto.PostResDto{
		AuditableResDto: dto.AuditableResDto{
			Id:        post.Id,
			CreatedAt: post.CreatedAt,
			DeletedAt: post.DeletedAt.Time,
			UpdatedAt: post.UpdatedAt,
		},
		Title:       post.Title,
		Description: post.Description,
		UserId:      post.UserId,
	}

	ctx.JSON(http.StatusOK, postRes)
}

func (handler *PostController) UpdatePost(ctx *httpContext.CustomContext) {
	postId := ctx.Param("id")
	user := ctx.GetUser()
	var reqDto postDto.UpdatePostReqDto
	if err := ctx.ShouldBindJSON(&reqDto); err != nil {
		ctx.Error(exception.NewUnprocessableEntityException(ctx.GetRequestId(), err))
		return
	}

	post, err := handler.postService.UpdatePost(user.Id, postId, &reqDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	postRes := postDto.PostResDto{
		AuditableResDto: dto.AuditableResDto{
			Id:        post.Id,
			CreatedAt: post.CreatedAt,
			DeletedAt: post.DeletedAt.Time,
			UpdatedAt: post.UpdatedAt,
		},
		Title:       post.Title,
		Description: post.Description,
		UserId:      post.UserId,
	}

	ctx.JSON(http.StatusOK, postRes)
}

func (handler *PostController) DeletePost(ctx *httpContext.CustomContext) {
	postId := ctx.Param("id")
	user := ctx.GetUser()

	err := handler.postService.SoftDeletePost(user.Id, postId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, true)
}
