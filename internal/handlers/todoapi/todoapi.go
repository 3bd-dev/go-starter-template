package todoapi

import (
	"encoding/json"
	"net/http"

	"github.com/3bd-dev/go-starter-template/internal/models"
	"github.com/3bd-dev/go-starter-template/internal/services"
)

type api struct {
	service *services.TodoService
}

func newapi(svc *services.TodoService) *api {
	return &api{service: svc}
}

func (h *api) Create(w http.ResponseWriter, r *http.Request) {
	var req models.Todo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": req.ID})
}

func (h *api) Update(w http.ResponseWriter, r *http.Request) {
	var req models.Todo
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.Update(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func (h *api) List(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}
