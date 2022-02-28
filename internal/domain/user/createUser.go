package user

import (
	entities "forum/internal/model"
)

func (s *Service) CreateUser(user entities.User) error {
	if err := CanRegister(user); err != nil {
		return err
	}
	err := s.Storage.AddUser(user.Name, user.Login, user.Password1)
	return err
}
