package post

import (
	entities "forum/internal/model"
)

func (r *PostRepository) GetCommLikes(idComm int, comment *entities.Comm) (*entities.Comm, error) {
	err := r.db.QueryRow(`SELECT COUNT (like)
	FROM CommAct
	WHERE commId = ?`, idComm).Scan(&comment.Likes)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
