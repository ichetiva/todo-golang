package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ichetiva/todo-golang/config"
	"github.com/ichetiva/todo-golang/internal/controllers/tasks"
)

func TaskRoute(router *mux.Router, cfg *config.Config) {
	controller := tasks.Controller{
		Config: cfg,
	}

	// Get task by id /api/v1/task?id=1
	router.HandleFunc(
		"/api/v1/task", controller.GetTask,
	).Methods(http.MethodGet)
	// Get all tasks
	router.HandleFunc(
		"/api/v1/task/all", controller.GetAllTasks,
	).Methods(http.MethodGet)
	// Get done tasks
	router.HandleFunc(
		"/api/v1/task/done", controller.GetDoneTasks,
	).Methods(http.MethodGet)
	// Get not done tasks
	router.HandleFunc(
		"/api/v1/task/notDone", controller.GetNotDoneTasks,
	).Methods(http.MethodGet)
	// Create task
	router.HandleFunc(
		"/api/v1/task/create", controller.AddTask,
	).Methods(http.MethodPost)
	// Mark task as done
	router.HandleFunc(
		"/api/v1/task/done", controller.DoneTask,
	).Methods(http.MethodPut)
	// Delete task
	router.HandleFunc(
		"/api/v1/task/delete", controller.DeleteTask,
	).Methods(http.MethodDelete)
}
