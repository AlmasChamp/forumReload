package post

import (
	entities "forum/internal/model"
)

func (r *PostRepository) GetAllComm(indPost int, post *entities.Post, commAct *entities.Comm) (*entities.Post, error) {
	rows, err := r.db.Query(`SELECT id,body
	FROM Comments
	WHERE postId = ?`, indPost)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&commAct.ID, &commAct.UserComm)

		r.db.QueryRow(`SELECT COUNT (like)
		FROM CommAct
		WHERE commId = ?`, commAct.ID).Scan(&commAct.Likes)

		r.db.QueryRow(`SELECT COUNT (dislike)
		FROM CommAct
		WHERE commId = ?`, commAct.ID).Scan(&commAct.Dislikes)
		// comment.ID = indPost
		post.AllComm = append(post.AllComm, *commAct)
	}
	return post, nil

}
