package post

func (s *PostService) AddPostComm(textPostComm string, UserId string, indPost int) error {
	if err := s.Storage.AddPostComm(textPostComm, UserId, indPost); err != nil {
		return err
	}
	return nil
}
