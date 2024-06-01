package service

import (
	"github.com/DLLenjoyer/todo-apiha/models"
    "github.com/DLLenjoyer/todo-apiha/repository"
)

type TodoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repository: repo}
}

func (s *TodoService) GetAll() []models.Todo {
	return s.repository.GetAll()
}

func (s *TodoService) GetByID(id string ) (*models.Todo, error) {
	return s.repository.GetByID(id)
}

func (s *TodoService) Add(todo *models.Todo) error {
	return s.repository.Add(todo)
}

func (s *TodoService) Update(todo *models.Todo) error {
	return s.repository.Update(todo)
}

func (s *TodoService) Delete(id string) error {
	return s.repository.Delete(id)
}
