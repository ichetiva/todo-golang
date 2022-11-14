package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ichetiva/todo-golang/config"
	"github.com/ichetiva/todo-golang/internal/controllers/todo"
)

func TodoRoute(router *mux.Router, cfg *config.Config) {
	controller := todo.Controller{
		Config: cfg,
	}

	// Get todo by id /api/v1/todo?id=1
	router.HandleFunc(
		"/api/v1/todo", controller.GetTodo,
	).Methods(http.MethodGet)
	// Get all todos
	router.HandleFunc(
		"/api/v1/todo/all", controller.GetAllTodos,
	).Methods(http.MethodGet)
	// Get done todos
	router.HandleFunc(
		"/api/v1/todo/done", controller.GetDoneTodos,
	).Methods(http.MethodGet)
	// Get not done todos
	router.HandleFunc(
		"/api/v1/todo/notDone", controller.GetNotDoneTodos,
	).Methods(http.MethodGet)
	// Create todo
	router.HandleFunc(
		"/api/v1/todo/create", controller.AddTodo,
	).Methods(http.MethodPost)
	// Mark todo as done
	router.HandleFunc(
		"/api/v1/todo/done", controller.DoneTodo,
	).Methods(http.MethodPut)
	// Delete todo
	router.HandleFunc(
		"/api/v1/todo/delete", controller.DeleteTodo,
	).Methods(http.MethodDelete)
}
