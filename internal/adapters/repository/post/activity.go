package post

func (r *PostRepository) Activity(UserId string, indPost int, DisOrLike string) int {
	activity := 0

	if DisOrLike == "like" {
		r.db.QueryRow(`SELECT like
		FROM DisLikes
		WHERE userId = ? AND postIdInComm = ?`, UserId, indPost).Scan(&activity)
		// fmt.Println(DisLike, UserId, indPost, "///////////////////////////////DeleOrAddlike")
		return activity
	}
	r.db.QueryRow(`SELECT dislike
	FROM DisLikes
	WHERE UserId = ? AND postIdInComm = ?`, UserId, indPost).Scan(&activity)
	// fmt.Println(DisLike, UserId, indPost, "///////////////////////////////DeleOrAddlike")
	return activity
}
