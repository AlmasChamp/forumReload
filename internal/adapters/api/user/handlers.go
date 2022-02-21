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
}
