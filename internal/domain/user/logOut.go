package user

import (
	"log"
	"net/http"
)

func (s *Service) LogOut(cookieVal string) *http.Cookie {

	cookie := &http.Cookie{}

	if err := s.Storage.DeleteCookie(cookieVal); err != nil {
		log.Println(err)
		return nil
	}

	cookie = &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	return cookie
}
