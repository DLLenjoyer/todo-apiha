package repository

import (
	"github.com/DLLenjoyer/todo-apiha/models"	
)

type TodoRepository interface {
	GetAll() []models.Todo
	GetByID(id string) (*models.Todo, error)
	Add(todo *models.Todo) error
	Update(todo *models.Todo) error
	Delete(id string) error
}

type InMemoryTodoRepository struct {
	Todos []models.Todo
}

func (r *InMemoryTodoRepository) GetAll() []models.Todo {
	return r.Todos
}

func (r *InMemoryTodoRepository) GetByID(id string) (*models.Todo, error) {
	for _, todo := range r.Todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, nil
}

func (r *InMemoryTodoRepository) Add(todo *models.Todo) error {
	r.Todos = append(r.Todos, *todo)
	return nil
}

func (r *InMemoryTodoRepository) Update(todo *models.Todo) error {
	for i, t := range r.Todos {
		if t.ID == todo.ID {
			r.Todos[i] = *todo
			return nil
		}
	}
	return nil
}

func (r *InMemoryTodoRepository) Delete(id string) error {
	for i, todo := range r.Todos {
		if todo.ID == id {
			r.Todos = append(r.Todos[:i], r.Todos[i+1:]...)
			return nil
		}
	}
	return nil
}
