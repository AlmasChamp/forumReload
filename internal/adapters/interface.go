package adapters

import (
	entities "forum/internal/model"
	"net/http"
)

type Service interface {
	CreateUser(input entities.User) error
	LogInUser(email, password string) (*http.Cookie, error)
	LogOut(cookieVal string) *http.Cookie
	AllPost() ([]entities.Post, error)
	ValueCookie(userCookie string) (string, error)
}

type Handler interface {
	Register(mux *http.ServeMux)
}
