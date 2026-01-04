package server

import (
	routes "GranmaCakesAPI/internal/routers"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router http.Handler
}

func New() *Server {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP) requires NGINX, Proxy, non-public project etc
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	routes.Register(r)

	return &Server{
		Router: r,
	}
}
