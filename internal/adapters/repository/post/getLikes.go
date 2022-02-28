package post

import entities "forum/internal/model"

func (r *PostRepository) GetLikes(indPost int, post *entities.Post) (*entities.Post, error) {

	err := r.db.QueryRow(`SELECT COUNT(like)
	FROM DisLikes
	WHERE PostIdInComm = ?`, indPost).Scan(&post.Likes)
	if err != nil {
		return nil, err
	}
	return post, nil
}
