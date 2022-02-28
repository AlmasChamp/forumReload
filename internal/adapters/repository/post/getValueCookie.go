package post

func (r *PostRepository) GetValueCookie(userCookie string) (string, error) {
	cookieFromDb := ""
	err := r.db.QueryRow(`SELECT value
	FROM Cookie
	WHERE value = ?`, userCookie).Scan(&cookieFromDb)
	if err != nil {
		return "", err
	}
	return cookieFromDb, nil
}
