package post

func (s *PostService) UserId(cookie string) (string, error) {
	res, err := s.Storage.GetUserId(cookie)
	if err != nil {
		return "", err
	}
	return res, nil
}
