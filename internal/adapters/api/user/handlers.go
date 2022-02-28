package user

import (
	"forum/internal/adapters"
	"forum/internal/app/middle"
	"net/http"
)

type Handler struct {
	Service adapters.Service
}

func NewHandler(service adapters.Service) adapters.Handler {
	return &Handler{Service: service}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/regPage", h.RegPage)
	mux.HandleFunc("/logPage", h.logPage)
	mux.HandleFunc("/", middle.MiddleWare(h.mainPage))
	fs := http.FileServer(http.Dir("./view/static"))
	mux.Handle("/view/static/", http.StripPrefix("/view/static/", fs))
	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
}
