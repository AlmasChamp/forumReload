package post

import (
	"forum/internal/app/middle"
	"net/http"
)

type PostHandler struct {
	Service PostService
}

func NewHandler(service PostService) Handler {
	return &PostHandler{Service: service}
}

func (h *PostHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/createPostPage", middle.MiddleWare(h.createPostPage))
	mux.HandleFunc("/postPage/", middle.MiddleWare(h.postPage))
	// mux.HandleFunc("/", middle.MiddleWare(h.mainPage))
}
