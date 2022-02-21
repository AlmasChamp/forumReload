package post

import (
	"database/sql"
	"forum/internal/domain/post"
)

type PostRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) post.PostStorage {
	return &PostRepository{db: db}
}
