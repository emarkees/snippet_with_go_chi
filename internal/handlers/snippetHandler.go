package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emarkees/chi/internal/app"
	"github.com/emarkees/chi/internal/models"
)

func Home(app *app.Application) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			notFound(w)
			return
		}

		snippets, err := app.Snippets.Latest(r.Context())
		if err != nil {
			serverError(app, w, err)
			return
		}

		data := newTemplateData(app, r)
		data.Snippets = snippets

		render(app, w, http.StatusOK, "home.tmpl", data)
	}
}

func CreateSnippet(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			clientError(w, http.StatusMethodNotAllowed)
			return
		}

		title := "O snail"
		content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
		expiresAt := 7

		id, err := app.Snippets.Insert(r.Context(), title, content, expiresAt)
		if err != nil {
			serverError(app, w, err)
			return
		}

		app.InfoLog.Printf("Snippet created with ID %d", id)
		http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

		// w.WriteHeader(http.StatusCreated)

		// w.Write([]byte("Snippet created successfully"))
	}
}

func ViewSnippet(app *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			notFound(w)
			return
		}

		snippet, err := app.Snippets.Get(r.Context(), id)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				notFound(w)
			} else {
				serverError(app, w, err)
			}
			return
		}

		data := newTemplateData(app, r)
		data.Snippet = snippet

		render(app, w, http.StatusOK, "view.tmpl", data)
	}
}
