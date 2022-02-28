package post

func (r *PostRepository) RemoveAct(UserId string, idComm int, commActivity string, num int) error {
	if num == 1 && commActivity == "like" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO CommAct (id,like,userId,commId)
		VALUES ((SELECT id
			FROM CommAct
			WHERE userId = $1 AND commId = $2),$3,$4,$5)`, UserId, idComm, nil, UserId, idComm)
		if err != nil {
			return err
		}
		return nil
	} else if num == 0 && commActivity == "like" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO CommAct (id,like,userId,commId)
		VALUES ((SELECT id
			FROM CommAct
			WHERE userId = $1 AND commId = $2),$3,$4,$5)`, UserId, idComm, 1, UserId, idComm)
		if err != nil {
			return err
		}
		return nil
	} else if num == 1 && commActivity == "dislike" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO CommAct (id,dislike,userId,commId)
		VALUES ((SELECT id
			FROM CommAct
			WHERE userId = $1 AND commId = $2),$3,$4,$5)`, UserId, idComm, nil, UserId, idComm)
		if err != nil {
			return err
		}
		return nil
	} else if num == 0 && commActivity == "dislike" {
		_, err := r.db.Exec(`INSERT OR REPLACE INTO CommAct (id,dislike,userId,commId)
		VALUES ((SELECT id
			FROM CommAct
			WHERE userId = $1 AND commId = $2),$3,$4,$5)`, UserId, idComm, 1, UserId, idComm)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
