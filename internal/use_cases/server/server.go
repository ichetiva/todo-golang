package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ichetiva/todo-golang/config"
	"github.com/ichetiva/todo-golang/internal/routes"
	"github.com/ichetiva/todo-golang/pkg/postgres"
	"github.com/ichetiva/todo-golang/pkg/postgres/models"
)

type Server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) Run(cfg *config.Config) error {
	router := s.GetRouter(cfg)

	server := &http.Server{
		Addr:         s.Addr,
		Handler:      router,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	db, err := postgres.GetDatabase(cfg)
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Todo{})

	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetRouter(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	routes.TodoRoute(router, cfg)
	return router
}
