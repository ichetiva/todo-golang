package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ichetiva/todo-golang/internal/controllers/todo"
)

func TodoRoute(router *mux.Router) {
	controller := todo.Controller{}

	router.HandleFunc(
		"/api/v1/todo", controller.GetTodo,
	).Methods(http.MethodGet)
}
