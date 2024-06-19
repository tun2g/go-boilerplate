package post

import (
	postDto "fist-app/src/apis/dto/post"
	"fist-app/src/apis/model"
	repository "fist-app/src/apis/repository/post"
	pageDto "fist-app/src/shared/dto"
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/shared/utils"
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

func (srv *postService) GetPostsByUserId(
	ctx *httpContext.CustomContext,
	userId string,
	dto *pageDto.PageOptionsDto,
)(*pageDto.PageDto, error){
	posts, err := srv.postRepository.GetPostsByUserIdAndPageDto(userId, dto)
	if(err != nil){
		return nil, err
	}

	count, err := srv.postRepository.CountByUserIdAndPageDto(userId, dto)
	if(err != nil){
		return nil, err
	}
	
	entities := make([]interface{}, len(*posts))
	for i, post := range *posts {
		entities[i] = post
	}

	pageRes := utils.GeneratePaginationResult(count, entities, dto)

	return pageRes, nil
}

func (srv *postService) SoftDeletePost(userId string, postId string) (error){
	err := srv.postRepository.SoftDeletePost(userId, postId)
	return err
}


func (srv *postService) UpdatePost(userId string, postId string, updatedPost *postDto.UpdatePostReqDto) (*model.Post, error){
	post, err := srv.postRepository.UpdatePost(userId, postId, updatedPost)
	return post, err
}


func (srv *postService) GetPost(userId string, postId string) (*model.Post, error){
	post, err := srv.postRepository.GetPost(userId, postId)
	return post, err
}
