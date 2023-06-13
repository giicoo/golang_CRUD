package app

import (
	"database/sql"

	"github.com/giicoo/golang_CRUD/internal/config"
	http_handlers "github.com/giicoo/golang_CRUD/internal/delivery/http"
	repository "github.com/giicoo/golang_CRUD/internal/repository/sqlite"
	"github.com/giicoo/golang_CRUD/internal/server"
	"github.com/giicoo/golang_CRUD/internal/service"
)

func Run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repository.NewRepo(db, cfg)
	service := service.NewServices(repo)
	delivery := http_handlers.NewHandler(service)

	r := delivery.Init()
	if err := repo.InitDB(); err != nil {
		return err
	}

	srv := server.NewServer(cfg.Host, r)
	if err := srv.Start(); err != nil {
		return err
	}
	return nil
}
