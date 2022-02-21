package post

import "net/http"

type PostService interface {
}

type Handler interface {
	Register(mux *http.ServeMux)
}
