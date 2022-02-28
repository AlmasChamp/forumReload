package post

func (r *PostRepository) CommAct(UserId string, idComm int, DisOrLike string) int {
	activity := 0

	if DisOrLike == "like" {
		r.db.QueryRow(`SELECT like
		FROM CommAct
		WHERE userId = ? AND commId = ?`, UserId, idComm).Scan(&activity)
		// fmt.Println(DisLike, UserId, idComm, "///////////////////////////////DeleOrAddlike")
		return activity
	}
	r.db.QueryRow(`SELECT dislike
	FROM CommAct
	WHERE UserId = ? AND commId = ?`, UserId, idComm).Scan(&activity)
	// fmt.Println(DisLike, UserId, indPost, "///////////////////////////////DeleOrAddlike")
	return activity
}
