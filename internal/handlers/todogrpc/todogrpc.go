package todogrpc

import (
	"context"

	"github.com/3bd-dev/go-starter-template/internal/models"
	"github.com/3bd-dev/go-starter-template/internal/services"
	pbtodo "github.com/3bd-dev/go-starter-template/rpc/todo"
)

type server struct {
	pbtodo.UnimplementedTodoServer
	service *services.TodoService
}

func newServer(svc *services.TodoService) *server {
	return &server{service: svc}
}

func (s *server) Create(ctx context.Context, req *pbtodo.CreateRequest) (*pbtodo.CreateResponse, error) {
	// Implement your gRPC service logic here
	todo := &models.Todo{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Completed:   req.GetCompleted(),
	}

	err := s.service.Create(todo)
	if err != nil {
		return nil, err
	}

	return &pbtodo.CreateResponse{
		Id: todo.ID, // Assuming the service generates and sets the ID
	}, nil
}

func (s *server) Update(ctx context.Context, req *pbtodo.UpdateRequest) (*pbtodo.UpdateResponse, error) {
	// Implement your gRPC service logic here
	todo := &models.Todo{
		ID:          req.GetId(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Completed:   req.GetCompleted(),
	}

	err := s.service.Update(todo)
	if err != nil {
		return nil, err
	}

	return &pbtodo.UpdateResponse{
		Success: true,
	}, nil
}

func (s *server) List(ctx context.Context, req *pbtodo.ListRequest) (*pbtodo.ListResponse, error) {
	// Implement your gRPC service logic here
	todos, err := s.service.List()
	if err != nil {
		return nil, err
	}

	var pbTodos []*pbtodo.Item
	for _, t := range todos {
		pbTodos = append(pbTodos, &pbtodo.Item{
			Id:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
		})
	}

	return &pbtodo.ListResponse{
		Items: pbTodos,
	}, nil
}
