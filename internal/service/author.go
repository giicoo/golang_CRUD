package service

import (
	"github.com/giicoo/golang_CRUD/internal/models"
	"github.com/giicoo/golang_CRUD/internal/repository"
)

type AuthorService struct {
	repo repository.Repo
}

func NewAuthorService(repo repository.Repo) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}

func (s *AuthorService) NewAuthor(name string) (int, error) {
	return s.repo.AddAuthor(name)
}

func (s *AuthorService) GetAuthor(au models.Author) ([]models.Author, error) {
	id := au.Author_id
	name := au.Name
	return s.repo.GetAuthor(id, name)
}

func (s *AuthorService) DeleteAuthor(id int) error {
	return s.repo.DeleteAuthor(id)
}
