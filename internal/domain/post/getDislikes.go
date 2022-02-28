package post

import (
	"errors"
	entities "forum/internal/model"
)

func (s *PostService) GetDislikes(indPost int, post *entities.Post) (*entities.Post, error) {
	if indPost < 0 {
		return nil, errors.New("Number less than 0")
	}
	post, err := s.Storage.GetDislikes(indPost, post)
	return post, err
}
