package post

import (
	commonDto "fist-app/src/shared/dto"
)

type PostResDto struct {
	commonDto.AuditableResDto
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"userId"`
}
