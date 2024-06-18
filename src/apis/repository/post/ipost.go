package post

import "fist-app/src/apis/model"

type PostRepository interface {
	StorePost(post *model.Post) (*model.Post, error)
}