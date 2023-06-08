package app

import (
	http_handlers "github.com/giicoo/golang_CRUD/internal/delivery/http"
	repository "github.com/giicoo/golang_CRUD/internal/repository/sqlite"
	"github.com/giicoo/golang_CRUD/internal/server"
	"github.com/giicoo/golang_CRUD/internal/service"
)

func Run() error {
	repo := repository.NewRepo()
	service := service.NewService(repo)
	delivery := http_handlers.NewHandler(service)

	delivery.Init()

	srv := server.NewServer("localhost:8888", delivery.Handler)
	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
