package repository

import "github.com/giicoo/golang_CRUD/internal/models"

type Repo interface {
	AddAuthor(name string) (int, error)
	GetAuthor(id int, name string) ([]models.Author, error)
	DeleteAuthor(id int) error
}
