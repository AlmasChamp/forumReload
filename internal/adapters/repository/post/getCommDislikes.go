package post

import (
	entities "forum/internal/model"
)

func (r *PostRepository) GetCommDislikes(idComm int, comment *entities.Comm) (*entities.Comm, error) {
	err := r.db.QueryRow(`SELECT COUNT (dislike)
	FROM CommAct
	// WHERE commId = ?`, idComm).Scan(&comment.Dislikes)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
