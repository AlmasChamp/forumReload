package post

func (r *PostRepository) ReplaceAct(UserId string, indPost int, disOrLike string, num int) error {

	if num == 1 && disOrLike == "like" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO DisLikes (id,like,userId,postIdInComm)
	VALUES ((SELECT id
		FROM DisLikes
		WHERE userId = $1 AND postIdInComm = $2),$3,$4,$5)`, UserId, indPost, nil, UserId, indPost)
		if err != nil {
			return err
		}
		return nil
	} else if num == 0 && disOrLike == "like" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO DisLikes (id,like,userId,postIdInComm)
	VALUES ((SELECT id
		FROM DisLikes
		WHERE userId = $1 AND postIdInComm = $2),$3,$4,$5)`, UserId, indPost, 1, UserId, indPost)
		if err != nil {
			return err
		}
		return nil
	} else if num == 1 && disOrLike == "dislike" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO DisLikes (id,dislike,userId,postIdInComm)
		VALUES ((SELECT id
			FROM DisLikes
			WHERE userId = $1 AND postIdInComm = $2),$3,$4,$5)`, UserId, indPost, nil, UserId, indPost)
		if err != nil {
			return err
		}
		return nil
	} else if num == 0 && disOrLike == "dislike" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO DisLikes (id,dislike,userId,postIdInComm)
		VALUES ((SELECT id
			FROM DisLikes
			WHERE userId = $1 AND postIdInComm = $2),$3,$4,$5)`, UserId, indPost, 1, UserId, indPost)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
