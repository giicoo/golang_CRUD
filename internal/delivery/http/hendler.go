package http_handlers

import (
	"fmt"
	"net/http"

	"github.com/giicoo/golang_CRUD/internal/service"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	services *service.Services
}

func NewHandler(service *service.Services) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) Init() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", h.index)
	router.POST("/add/author", h.addAuthor)
	router.GET("/get/author", h.getAuthor)
	router.POST("/del/author", h.deleteAuthor)
	return router
}

func (h *Handler) index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Hello, World")
}
