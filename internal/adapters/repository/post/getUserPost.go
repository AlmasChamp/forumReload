package post

import (
	"fmt"
	entities "forum/internal/model"
)

func (r *PostRepository) GetUserPost(post *entities.Post, indPost int) (*entities.Post, error) {

	rows, err := r.db.Query(`SELECT title, body
	FROM Posts
	WHERE id = ?`, indPost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&post.Title, &post.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return post, nil
}
