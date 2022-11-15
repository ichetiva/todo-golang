package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ichetiva/todo-golang/config"
	_ "github.com/ichetiva/todo-golang/docs"
	"github.com/ichetiva/todo-golang/internal/routes"
	"github.com/ichetiva/todo-golang/pkg/postgres"
	"github.com/ichetiva/todo-golang/pkg/postgres/models"
	httpSwagger "github.com/swaggo/http-swagger"
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
	db.AutoMigrate(&models.Task{})

	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetRouter(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/swagger/").Handler(
		httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", cfg.App.Host, cfg.App.Port)),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		),
	).Methods(http.MethodGet)
	routes.TaskRoute(router, cfg)
	return router
}
