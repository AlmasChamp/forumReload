package user

import (
	"context"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func (s *Service) LogInUser(email, password string) (*http.Cookie, error) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	passwordFromDb, err := s.Storage.GetUserPassword(email)
	if err != nil || passwordFromDb != password {
		log.Println(err)
		return nil, err
	}
	id, err := s.Storage.GetUserId(email)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cookie := &http.Cookie{
		Name:   "session",
		Value:  uuid.NewV1().String(),
		MaxAge: 200,
	}

	if err := s.Storage.SetCookie(cookie.Value, cookie.MaxAge, id); err != nil {
		log.Println(err)
		return nil, err
	}

	return cookie, nil
}
