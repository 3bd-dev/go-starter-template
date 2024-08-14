package todoapi

import (
	"net/http"

	"github.com/3bd-dev/go-starter-template/internal/services"
	"github.com/gorilla/mux"
)

// Routes adds specific routes for this group.
func Routes(router *mux.Router, svc *services.TodoService) {
	api := newapi(svc)
	todos := router.PathPrefix("/todos").Subrouter()
	todos.HandleFunc("/", api.List).Methods(http.MethodGet)    // List all todos
	todos.HandleFunc("/", api.Create).Methods(http.MethodPost) // Create a new todo
	todos.HandleFunc("/{id}", api.Update).Methods(http.MethodPut)
}
