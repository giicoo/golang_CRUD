package service

import (
	repo "github.com/giicoo/golang_CRUD/internal/repository/sqlite"
)

type Service struct {
	repo *repo.Sqlite
}

func NewService(r *repo.Sqlite) *Service {
	return &Service{
		repo: r,
	}
}
