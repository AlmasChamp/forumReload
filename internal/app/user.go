package app

import (
	"database/sql"
	"forum/internal/adapters"
	userApi "forum/internal/adapters/api/user"
	userRep "forum/internal/adapters/repository/user"
	"forum/internal/domain"
	"forum/internal/domain/user"
)
//go:generate mockgen -source=user.go -destination=mocks/mock.go
type UserComposites struct {
	Storage domain.Storage
	Service adapters.Service
	Handler adapters.Handler
}

func NewUserComposites(db *sql.DB) (*UserComposites, error) {
	// Init Storage
	repository := userRep.NewRepository(db)
	// Init Service
	service := user.NewService(repository)
	// Init Handler
	handler := userApi.NewHandler(service)

	return &UserComposites{
		Storage: repository,
		Service: service,
		Handler: handler,
	}, nil
}
