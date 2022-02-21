package user

import (
	entities "forum/internal/model"
	"log"
)

func (s *Service) AllPost() ([]entities.Post, error) {
	posts, err := s.Storage.GetAllPosts()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return posts, nil
}
