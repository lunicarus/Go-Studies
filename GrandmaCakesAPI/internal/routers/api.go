package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"GranmaCakesAPI/internal/handlers"
	"GranmaCakesAPI/internal/middleware"
)

func Register(r chi.Router) {
	r.Route("/api", func(r chi.Router) {

		r.Post("/login", handlers.Login)

		r.Route("/clients", func(r chi.Router) {
			//public
			r.Get("/", handlers.ListClients)

			//protected
			r.Group(func(r chi.Router) {
				r.Use(middleware.Auth)
				r.Get("/{id}", handlers.GetClient)
				r.Post("/", handlers.CreateClient)
				r.Put("/{id}", handlers.UpdateClient)
				r.Delete("/{id}", handlers.DeleteClient)
			})
		})

		r.Route("/cakes", func(r chi.Router) {

			r.Get("/", handlers.ListCakes)

			r.Group(func(r chi.Router) {
				r.Use(middleware.Auth)
				//TODO: In a perfect world, I would have time to build  FULL REST API grandma Bakery Backend
				//! But this is not a perfect world...
				// r.Post("/", handlers.CreateCake)
				// r.Put("/{id}", handlers.UpdateCake)
				// r.Delete("/{id}", handlers.DeleteCake)
				r.Get("/{id}", handlers.GetCake)
			})
		})
	})

	// 404 fallback
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "route not found", http.StatusNotFound)
	})
}
