package post

type CreatePostReqDto struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required,min=8"`
}
