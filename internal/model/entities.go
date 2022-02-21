package entities

import "net/http"

type Server struct {
	httpServer *http.Server
}

type User struct {
	Name       string
	Login      string
	Password1  string
	Password2  string
	Cookie     string
	Categories []string
	Islogged   bool
	Error
}

type Post struct {
	ID       string
	Title    string
	Body     string
	Likes    int
	Dislikes int
	AllComm  []Comm
}

type Comm struct {
	ID       int
	UserComm string
	Likes    int
	Dislikes int
}

type Out struct {
	User
	UsersPosts []Post
}

type Error struct {
	MSG string
	Err bool
}
