package post

func (r *PostService) RemoveAct(UserId string, idComm int, commActivity string, num int) error {
	if err := r.Storage.RemoveAct(UserId, idComm, commActivity, num); err != nil {
		return err
	}
	return nil
}
