package user

import (
	"forum/internal/adapters"
	"forum/internal/domain"
)

type Service struct {
	Storage domain.Storage
}

func NewService(repository domain.Storage) adapters.Service {
	return &Service{Storage: repository}
}
