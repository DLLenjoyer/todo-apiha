package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/DLLenjoyer/todo-apiha/models"
    "github.com/DLLenjoyer/todo-apiha/service"
	"github.com/gorilla/mux"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/todos", h.getAllTodos).Methods(http.MethodGet)
	r.HandleFunc("/todos", h.addTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", h.updateTodo).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", h.deleteTodo).Methods(http.MethodDelete)
}

func (h *TodoHandler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := h.service.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) addTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	if err := h.service.Add(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) updateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	if err := h.service.Update(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
