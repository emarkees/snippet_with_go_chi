package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/emarkees/chi/internal/handlers"
)

func SetUpRoutes(r *chi.Mux) {
	r.Get("/", handlers.Home)
	r.Post("/snippet/create", handlers.CreateSnippet)
	r.Get("/snippet/view", handlers.ViewSnippet)
}