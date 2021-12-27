package cofe_api

import (
	"net/http"
	"time"

	// third party

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// my own

	"github.com/cofeGB/coffeGBBackend/internal/cofe_services"
)

var ServerTimeout = 60 * time.Second

func NewCofeAPIServer(addr, logLevel string, service cofe_services.CofeServices) (srv *http.Server) {
	// handler
	handler := CofeHandler{service: service}
	// mux
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(ServerTimeout))

	// endpoints
	router.Get("/api", handler.Hello)

	////http://localhost:3000/api/navMenu/
	router.Get("/api/navMenu", handler.GetListNawMenu)

	// http://localhost:3000/api/categories/
	router.Get("/api/categories", handler.GetCategories)

	return &http.Server{Addr: addr, Handler: router}
}
