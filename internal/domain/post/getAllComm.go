package post

import (
	"errors"
	entities "forum/internal/model"
)

func (s *PostService) GetAllComm(indPost int, post *entities.Post, comment *entities.Comm) (*entities.Post, error) {
	if indPost < 0 {
		return nil, errors.New("idPost les than 0")
	}
	post, err := s.Storage.GetAllComm(indPost, post, comment)
	return post, err
}
