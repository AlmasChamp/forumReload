package post

func (r *PostRepository) AddPostComm(textPostComm string, UserId string, indPost int) error {
	_, err := r.db.Exec(`INSERT INTO Comments (body,userId,postId)
			VALUES ($1,$2,$3)`, textPostComm, UserId, indPost)
	if err != nil {
		return err
	}
	return nil
}
