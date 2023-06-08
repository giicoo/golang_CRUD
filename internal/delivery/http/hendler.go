package http_handlers

import (
	"fmt"
	"net/http"

	"github.com/giicoo/golang_CRUD/internal/service"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Handler *httprouter.Router
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Handler: httprouter.New(),
		service: service,
	}
}

func (h *Handler) Init() {
	h.Handler.GET("/", h.index)
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Hello, World")
}
