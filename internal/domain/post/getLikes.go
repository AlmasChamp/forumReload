package post

import (
	"errors"
	entities "forum/internal/model"
)

func (s *PostService) GetLikes(indPost int, post *entities.Post) (*entities.Post, error) {
	if indPost < 0 {
		return nil, errors.New("Number less than 0")
	}
	post, err := s.Storage.GetLikes(indPost, post)
	return post, err
}
