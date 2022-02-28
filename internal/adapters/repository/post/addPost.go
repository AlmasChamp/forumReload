package post

func (r *PostRepository) AddPost(title, body, userId string) error {
	_, err := r.db.Exec(`INSERT INTO Posts (title,body,userId)
	VALUES($1,$2,$3)`, title, body, userId)
	if err != nil {
		return err
	}
	return nil
}
