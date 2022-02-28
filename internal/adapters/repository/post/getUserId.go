package post

func (r *PostRepository) GetUserId(userCookie string) (string, error) {
	userIdFromDb := ""
	err := r.db.QueryRow(`SELECT userId
	FROM Cookie
	WHERE value = ?`, userCookie).Scan(&userIdFromDb)
	if err != nil {
		return "", err
	}
	return userIdFromDb, nil
}
