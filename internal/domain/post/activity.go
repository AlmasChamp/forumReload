package post

func (s *PostService) Activity(UserId string, indPost int, DisOrLike string) int {
	activity := s.Storage.Activity(UserId, indPost, DisOrLike)
	return activity
}
