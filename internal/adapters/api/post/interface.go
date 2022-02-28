package post

import (
	entities "forum/internal/model"
	"net/http"
)

type PostService interface {
	ValueCookie(userCookie string) (string, error)
	UserId(string) (string, error)
	CreatePost(title, body, userId string) error
	GetPost(post *entities.Post, indPost int) (*entities.Post, error)
	GetLikes(indPost int, post *entities.Post) (*entities.Post, error)
	GetDislikes(indPost int, post *entities.Post) (*entities.Post, error)
	GetCommLikes(idComm int, comment *entities.Comm) (*entities.Comm, error)
	GetCommDislikes(idComm int, comment *entities.Comm) (*entities.Comm, error)
	GetAllComm(indPost int, post *entities.Post, commAct *entities.Comm) (*entities.Post, error)
	CommAct(UserId string, idComm int, DisOrLike string) int
	RemoveAct(UserId string, idComm int, commActivity string, num int) error
	Activity(UserId string, indPost int, disOrLike string) int
	ReplaceAct(UserId string, indPost int, disOrLike string, num int) error
	AddPostComm(textPostComm string, UserId string, indPost int) error
}

type Handler interface {
	Register(mux *http.ServeMux)
}
