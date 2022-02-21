package user

import (
	"context"
	"forum/internal/adapters"
	"forum/internal/domain"
	entities "forum/internal/model"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Service struct {
	Storage domain.Storage
}

func NewService(repository domain.Storage) adapters.Service {
	return &Service{Storage: repository}
}

func (s *Service) ValueCookie(userCookie string) (string, error) {
	cookieVal, err := s.Storage.GetValueCookie(userCookie)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return cookieVal, nil

}

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

func (s *Service) CreateUser(user entities.User) error {

	if err := CanRegister(user); err != nil {
		return err
	}

	err := s.Storage.AddUser(user.Name, user.Login, user.Password1)
	return err
}

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

func (s *Service) AllPost() ([]entities.Post, error) {
	posts, err := s.Storage.GetAllPosts()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return posts, nil
}
