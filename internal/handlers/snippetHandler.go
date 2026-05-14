package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/emarkees/chi/internal/app"
)

func Home(app *app.Application) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			notFound(w)
			return
		}

		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			"./ui/html/pages/home.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			serverError(app, w, err)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			serverError(app, w, err)
			return
		}
	}
}

func ViewSnippet(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			notFound(w)
			return
		}

		fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	}
}

func CreateSnippet(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allowed", http.MethodPost)
			clientError(w, http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Create a snippet"))
	}
}
