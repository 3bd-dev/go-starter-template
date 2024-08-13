package todoapi

import (
	"net/http"

	"github.com/3bd-dev/go-starter-template/internal/services"
)

// Routes adds specific routes for this group.
func Routes(mux *http.ServeMux, svc *services.TodoService) {
	api := newapi(svc)
	mux.HandleFunc("/todos/create", api.Create)
	mux.HandleFunc("/todos/update", api.Update)
	mux.HandleFunc("/todos/list", api.List)
}
