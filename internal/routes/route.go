package routes

import (
	"github.com/go-chi/chi/v5"
	// "github.com/emarkees/chi/internal/handlers"
)

func SetUpRoutes(r *chi.Mux, app *Application) {
	r.Get("/", Home(app))
	r.Post("/snippet/create", app.CreateSnippet)
	r.Get("/snippet/view", app.ViewSnippet)
}