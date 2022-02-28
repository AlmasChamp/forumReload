package post

import (
	"errors"
	entities "forum/internal/model"
)

func (s *PostService) GetCommDislikes(idComm int, comment *entities.Comm) (*entities.Comm, error) {
	if idComm < 0 {
		return nil, errors.New("idComm les than 0")
	}
	comm, err := s.Storage.GetCommLikes(idComm, comment)
	return comm, err
}
