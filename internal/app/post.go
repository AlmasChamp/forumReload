package app

import (
	"database/sql"
	handler "forum/internal/adapters/api/post"
	postApi "forum/internal/adapters/api/post"
	postRep "forum/internal/adapters/repository/post"
	"forum/internal/domain/post"
)

//go:generate mockgen -source=user.go -destination=mocks/mock.go
type PostComposites struct {
	Storage post.PostStorage
	Service postApi.PostService
	Handler handler.Handler
}

func NewPostComposites(db *sql.DB) (*PostComposites, error) {
	// Init Storage
	repository := postRep.NewRepository(db)
	// Init Service
	service := post.NewService(repository)
	// Init Handler
	handler := postApi.NewHandler(service)

	return &PostComposites{
		Storage: repository,
		Service: service,
		Handler: handler,
	}, nil
}
