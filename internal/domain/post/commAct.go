package post

func (s *PostService) CommAct(UserId string, idComm int, DisOrLike string) int {
	activity := s.Storage.CommAct(UserId, idComm, DisOrLike)
	return activity
}
