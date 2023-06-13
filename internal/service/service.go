package service

import (
	"github.com/giicoo/golang_CRUD/internal/models"
	"github.com/giicoo/golang_CRUD/internal/repository"
)

//go:generate $GOPATH/bin/mockgen -source=service.go -destination=mocks/mock.go
type AuthorServices interface {
	NewAuthor(name string) (int, error)
	GetAuthor(au models.Author) ([]models.Author, error)
	DeleteAuthor(id int) error
}

type Services struct {
	AuthorServices
}

func NewServices(r repository.Repo) *Services {
	return &Services{
		AuthorServices: NewAuthorService(r),
	}
}
