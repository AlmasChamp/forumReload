package post

func (s *PostService) CreatePost(title, body, userId string) error {
	if err := s.Storage.AddPost(title, body, userId); err != nil {
		return err
	}
	return nil
}
