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
	todos []models.Todo
}

func (r *InMemoryTodo) GetAll() []models.Todo {
	return r.todos
}

func (r *InMemoryTodo) GetByID(id string) (*models.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, nil
}

func (r *InMemoryTodo) Add(todo *models.Todo) error {
	r.todos = append(r.todos, *todo)
	return nil
}

func (r *InMemoryTodo) Update(todo *models.Todo) error {
	for i, t := range r.todos {
		if t.ID == todo.ID {
			r.todos[i] = *todo
			return nil
		}
	}
	return nil
}

func (r *InMemoryTodo) Delete(id string) error {
	for i, todo := range r.todos {
		if todo.ID == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return nil
}
