package server

import (
	"fmt"
	"log"
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
	log.Println("Preparing to start")

	router := s.GetRouter(cfg)

	server := &http.Server{
		Addr:         s.Addr,
		Handler:      router,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	log.Println("Run migrations")
	db, err := postgres.GetDatabase(cfg)
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})

	log.Printf("Running on http://%s:%s\n", cfg.App.Host, cfg.App.Port)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) GetRouter(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	// Swagger setting
	router.PathPrefix("/swagger/").Handler(
		httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", cfg.App.Host, cfg.App.Port)),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		),
	).Methods(http.MethodGet)

	// Redirect to swagger
	router.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/swagger/index.html", http.StatusSeeOther)
		},
	)

	// Tasks route
	routes.TaskRoute(router, cfg)
	// User route
	routes.UserRoute(router, cfg)

	return router
}
