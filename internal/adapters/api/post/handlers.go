package post

import (
	"net/http"
)

type PostHandler struct {
	Service PostService
}

func NewHandler(service PostService) Handler {
	return &PostHandler{Service: service}
}

func (h *PostHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/createPostPage", h.createPostPage)

	// mux.HandleFunc("/", middle.MiddleWare(h.mainPage))
}
