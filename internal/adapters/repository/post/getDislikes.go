package post

import entities "forum/internal/model"

func (r *PostRepository) GetDislikes(indPost int, post *entities.Post) (*entities.Post, error) {

	err := r.db.QueryRow(`SELECT COUNT(dislike)
	FROM DisLikes
	WHERE PostIdInComm = ?`, indPost).Scan(&post.Dislikes)
	if err != nil {
		return nil, err
	}
	return post, nil
}
