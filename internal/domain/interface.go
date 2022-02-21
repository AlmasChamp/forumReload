package domain

import entities "forum/internal/model"

//go:generate mockgen -source=interface.go -destination=mocks/mock.go

type Storage interface {
	AddUser(name, email, password string) error
	GetUserPassword(email string) (string, error)
	GetUserId(email string) (string, error)
	SetCookie(cookieVal string, cookieExp int, id string) error
	DeleteCookie(cookie string) error
	GetAllPosts() ([]entities.Post, error)
	GetValueCookie(userCookie string) (string, error)
}

