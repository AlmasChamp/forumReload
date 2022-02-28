package post

import (
	entities "forum/internal/model"
)

func (s *PostService) GetPost(post *entities.Post, indPost int) (*entities.Post, error) {
	post, err := s.Storage.GetUserPost(post, indPost)
	if err != nil {
		return post, err
	}
	return post, nil
}
