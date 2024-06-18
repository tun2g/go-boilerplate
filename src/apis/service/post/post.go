package post

import (
	"fist-app/src/apis/model"
	repository "fist-app/src/apis/repository/post"
	httpContext "fist-app/src/shared/http-context"
	postDto "fist-app/src/apis/dto/post"
)

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) *postService {
	return &postService{
		postRepository: postRepository,
	}
}

func (srv *postService) CreateNewPost(
	ctx *httpContext.CustomContext,
	userId string, 
	dto postDto.CreatePostReqDto,
) (*model.Post, error){
	var err error
	post, err := srv.postRepository.StorePost(&model.Post{
		UserId: userId,
		Title: dto.Title,
		Description: dto.Description,
	})
	return post, err;
}


