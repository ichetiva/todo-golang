package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ichetiva/todo-golang/internal/routes"
)

type Server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) Run() error {
	router := s.GetRouter()

	server := &http.Server{
		Addr:         s.Addr,
		Handler:      router,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetRouter() *mux.Router {
	router := mux.NewRouter()
	routes.TodoRoute(router)
	return router
}
