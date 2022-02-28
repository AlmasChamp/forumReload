package post

import "log"

func (s *PostService) ValueCookie(userCookie string) (string, error) {
	cookieVal, err := s.Storage.GetValueCookie(userCookie)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return cookieVal, nil

}
