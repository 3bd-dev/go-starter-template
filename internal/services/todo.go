package services

import (
	"errors"

	"github.com/3bd-dev/go-starter-template/internal/models"
)

type TodoRepo interface {
	List() ([]*models.Todo, error)
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
}

type TodoService struct {
	repo TodoRepo
}

func NewTodoService(repo TodoRepo) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	return s.repo.Create(todo)
}

func (s *TodoService) Update(todo *models.Todo) error {
	if todo.ID == "" {
		return errors.New("id is required")
	}

	return s.repo.Update(todo)
}

func (s *TodoService) List() ([]*models.Todo, error) {
	return s.repo.List()
}
