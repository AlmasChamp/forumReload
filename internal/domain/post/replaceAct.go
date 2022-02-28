package post

func (r *PostService) ReplaceAct(UserId string, indPost int, disOrLike string, num int) error {
	if err := r.Storage.ReplaceAct(UserId, indPost, disOrLike, num); err != nil {
		return err
	}
	return nil
}
