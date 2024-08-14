package inmemory

import (
	"errors"
	"sync"

	"github.com/3bd-dev/go-starter-template/internal/models"
	"github.com/google/uuid"
)

type TodoRepository struct {
	mu    sync.Mutex
	todos map[string]*models.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos: make(map[string]*models.Todo),
	}
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.ID = uuid.New().String() // Generate a unique ID
	r.todos[todo.ID] = todo
	return nil
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}

	r.todos[todo.ID] = todo
	return nil
}

func (r *TodoRepository) List() ([]*models.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*models.Todo
	for _, todo := range r.todos {
		list = append(list, todo)
	}

	return list, nil
}
