package post

import (
	"forum/internal/adapters/api/post"
)

type PostService struct {
	Storage PostStorage
}

func NewService(repository PostStorage) post.PostService {
	return &PostService{Storage: repository}
}
