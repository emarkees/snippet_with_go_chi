package routes

import (
	"net/http"

	"github.com/emarkees/chi/internal/app"
	"github.com/emarkees/chi/internal/handlers"
	"github.com/emarkees/chi/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.SecureHeaders)

	// File is serve through the http.FileServer
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	
	r.Get("/", handlers.Home(app))
	r.Post("/snippet/create", handlers.CreateSnippet(app))
	r.Get("/snippet/view", handlers.ViewSnippet(app))

	return r
}