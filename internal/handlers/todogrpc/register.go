package todogrpc

import (
	"github.com/3bd-dev/go-starter-template/internal/services"
	pbtodo "github.com/3bd-dev/go-starter-template/rpc/todo"
	"google.golang.org/grpc"
)

func Register(serv *grpc.Server, svc *services.TodoService) {
	s := newServer(svc)
	pbtodo.RegisterTodoServer(serv, s)
}
