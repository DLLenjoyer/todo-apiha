package main

import (
	"log"
	"net/http"

	"github.com/DLLenjoyer/todo-apiha/handlers"
	"github.com/DLLenjoyer/todo-apiha/models"
	"github.com/DLLenjoyer/todo-apiha/repository"
	"github.com/DLLenjoyer/todo-apiha/service"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.InMemoryTodoRepository{
		Todos: []models.Todo{},
	}

	todoService := service.NewTodoService(&repo)

	todoHandler := handlers.NewTodoHandler(todoService)

	r := mux.NewRouter()

	todoHandler.SetupRoutes(r)

	log.Println("starting server on :6030")
	log.Fatal(http.ListenAndServe(":6030", r))
}
